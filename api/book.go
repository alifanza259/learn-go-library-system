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

	// Create arguments for BorrowTx function
	accessTokenPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	borrowDetailArg := db.CreateBorrowParams{
		BookID:     req.ID,
		BorrowedAt: time.Unix(int64(req.BorrowDate), 0),
		ReturnedAt: time.Unix(int64(req.ReturnDate), 0),
	}
	createTransactionArg := db.CreateTransactionParams{
		MemberID: uuid.MustParse(accessTokenPayload.ID),
		Purpose:  "borrow",
		Status:   db.StatusPending,
		BorrowID: uuid.Nil,
	}

	transaction, err := server.db.BorrowTx(c, db.BorrowTxParams{
		CreateBorrowParams:      borrowDetailArg,
		CreateTransactionParams: createTransactionArg,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		if err == errors.New("books are not available to borrow") {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

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

type ListBorrowRequestParams struct {
	Status db.Status `form:"status"`
}

func (server *Server) listBorrowRequests(c *gin.Context) {
	var req ListBorrowRequestParams
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	accessTokenPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	histories, err := server.db.GetBorrowHistory(c, db.GetBorrowHistoryParams{
		MemberID: uuid.MustParse(accessTokenPayload.ID),
		Status: db.NullStatus{
			Status: req.Status,
			Valid:  req.Status != "",
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, histories)
}
