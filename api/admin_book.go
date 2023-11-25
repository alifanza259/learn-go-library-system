package api

import (
	"errors"
	"net/http"
	"time"

	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
	"github.com/alifanza259/learn-go-library-system/token"
	"github.com/alifanza259/learn-go-library-system/worker"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgtype"
)

type createBookRequest struct {
	Isbn        string `form:"isbn" binding:"required"`
	Title       string `form:"title" binding:"required"`
	Description string `form:"description" binding:"required"`
	Author      string `form:"author" binding:"required"`
	Genre       string `form:"genre" binding:"required"`
	Quantity    int    `form:"quantity" binding:"required"`
	PublishedAt int64  `form:"published_at" binding:"required"`
}

func (server *Server) createBook(c *gin.Context) {
	var req createBookRequest
	if err := c.ShouldBindWith(&req, binding.FormMultipart); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			c.JSON(http.StatusBadRequest, errorResponse(errors.New("image is not provided")))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	location, err := server.external.UploadAttachment(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateBookParams{
		Isbn:  req.Isbn,
		Title: req.Title,
		Description: pgtype.Text{
			String: req.Description,
			Valid:  req.Description != "",
		},
		Author: req.Author,
		ImageUrl: pgtype.Text{
			String: location,
			Valid:  true,
		},
		Genre:       req.Genre,
		Quantity:    int32(req.Quantity),
		PublishedAt: time.Unix(req.PublishedAt, 0),
	}

	book, err := server.db.CreateBook(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, book)
}

type UpdateBookRequestJSON struct {
	Isbn        string `form:"isbn"`
	Title       string `form:"title"`
	Description string `form:"description"`
	Author      string `form:"author"`
	Genre       string `form:"genre"`
	Quantity    *int   `form:"quantity"`
	PublishedAt *int64 `form:"published_at"`
}

type UpdateBookRequestURI struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) updateBook(c *gin.Context) {
	var reqForm UpdateBookRequestJSON
	var reqURI UpdateBookRequestURI

	if err := c.ShouldBindWith(&reqForm, binding.FormMultipart); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := c.ShouldBindUri(&reqURI); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Validate entry is exist
	_, err := server.db.GetBook(c, int32(reqURI.ID))
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// If update image, run upload to S3 process
	file, err := c.FormFile("image")
	if err != nil {
		if err != http.ErrMissingFile {
			c.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	var location string
	if file != nil {
		location, err = server.external.UploadAttachment(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	// Update entry
	var quantity int32
	var publishedAt time.Time
	if reqForm.PublishedAt != nil {
		publishedAt = time.Unix(*reqForm.PublishedAt, 0)
	}
	if reqForm.Quantity != nil {
		quantity = int32(*reqForm.Quantity)
	}

	arg := db.UpdateBookParams{
		ID: reqURI.ID,
		Isbn: pgtype.Text{
			String: reqForm.Isbn,
			Valid:  reqForm.Isbn != "",
		},
		Title: pgtype.Text{
			String: reqForm.Title,
			Valid:  reqForm.Title != "",
		},
		Description: pgtype.Text{
			String: reqForm.Description,
			Valid:  reqForm.Description != "",
		},
		Author: pgtype.Text{
			String: reqForm.Author,
			Valid:  reqForm.Author != "",
		},
		ImageUrl: pgtype.Text{
			String: location,
			Valid:  file != nil,
		},
		Genre: pgtype.Text{
			String: reqForm.Genre,
			Valid:  reqForm.Genre != "",
		},
		Quantity: pgtype.Int4{
			Valid: reqForm.Quantity != nil,
			Int32: quantity,
		},
		PublishedAt: pgtype.Timestamptz{
			Valid: reqForm.PublishedAt != nil,
			Time:  publishedAt,
		},
	}
	updatedBook, err := server.db.UpdateBook(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, updatedBook)
}

type DeleteBookRequest struct {
	ID int `uri:"id" binding:"required"`
}

func (server *Server) deleteBook(c *gin.Context) {
	var req DeleteBookRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Delete entry
	err := server.db.DeleteBook(c, int32(req.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusNoContent, "")
}

type ProcessBorrowReqRequest struct {
	ID     uuid.UUID `json:"transaction_id" binding:"required"`
	Status string    `json:"status" binding:"required"`
	Note   string    `json:"note"`
}

func (server *Server) processBorrowReq(c *gin.Context) {
	var req ProcessBorrowReqRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Get Borrow/Return Request
	transaction, err := server.db.GetTransactionAssociatedDetail(c, req.ID)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Update request status
	accessTokenPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.UpdateTransactionParams{
		AdminID: pgtype.UUID{
			Valid: true,
			Bytes: uuid.MustParse(accessTokenPayload.ID),
		},
		Status: db.Status(req.Status),
		Note: pgtype.Text{
			String: req.Note,
			Valid:  req.Note != "",
		},
		ID: req.ID,
	}
	updatedTrx, err := server.db.ProcessBorrowTx(c, db.ProcessBorrowTxParams{
		UpdateTransactionParams: arg,
		Transaction:             transaction,
		AfterUpdate: func(transaction db.GetTransactionAssociatedDetailRow, status db.Status, note string) error {
			return server.taskDistributor.DistributeTaskSendBorrowProcessedEmail(c, &worker.PayloadSendBorrowProcessedEmail{
				TransactionID: transaction.TrxID.String(),
				MemberEmail:   transaction.Email,
				MemberName:    transaction.FirstName,
				BookTitle:     transaction.BTitle,
				Status:        status,
				Note:          note,
			})
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, updatedTrx)
}
