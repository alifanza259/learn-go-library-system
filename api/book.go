package api

import (
	"errors"
	"net/http"
	"time"

	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
	"github.com/alifanza259/learn-go-library-system/token"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (server *Server) listBooks(c *gin.Context) {
	books, err := server.db.ListBooks(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, books)
}

type GetBookRequest struct {
	ID int `uri:"id" binding:"required"`
}

func (server *Server) getBook(c *gin.Context) {
	var req GetBookRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	book, err := server.db.GetBook(c, int32(req.ID))
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, book)
}

type BorrowBooksRequest struct {
	ID         int32 `json:"id" binding:"required"`
	BorrowDate int   `json:"borrow_date" binding:"required"`
	ReturnDate int   `json:"return_date" binding:"required"`
}

func (server *Server) borrowBooks(c *gin.Context) {
	var req BorrowBooksRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Find books in books table, check quantity
	book, err := server.db.GetBook(c, req.ID)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if book.Quantity == 0 {
		c.JSON(http.StatusBadRequest, errorResponse(errors.New("books are not available to borrow")))
		return
	}

	// TODO: Create db transaction
	// Create entry in borrow_details table
	borrowDetailArg := db.CreateBorrowParams{
		BookID:     req.ID,
		BorrowedAt: time.Unix(int64(req.BorrowDate), 0),
		ReturnedAt: time.Unix(int64(req.ReturnDate), 0),
	}
	borrowDetail, err := server.db.CreateBorrow(c, borrowDetailArg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Create entry in transactions table
	accessTokenPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.CreateTransactionParams{
		MemberID: uuid.MustParse(accessTokenPayload.ID),
		Purpose:  "borrow",
		Status:   db.StatusPending,
		BorrowID: borrowDetail.ID,
	}
	transaction, err := server.db.CreateTransaction(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, transaction)
}

type ReturnBooksRequest struct {
	ID uuid.UUID `json:"id" binding:"required"`
}

func (server *Server) returnBooks(c *gin.Context) {
	var req ReturnBooksRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Find the transaction and borrow_detail
	borrowTransaction, err := server.db.GetTransactionAndBorrowDetail(c, req.ID)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Create return transaction
	arg := db.CreateTransactionParams{
		BorrowID: borrowTransaction.BdID,
		MemberID: borrowTransaction.TrxMemberID,
		Purpose:  db.PurposeReturn,
		Status:   db.StatusPending,
	}
	returnTransaction, err := server.db.CreateTransaction(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, returnTransaction)
}