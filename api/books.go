package api

import (
	"net/http"
	"time"

	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
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

type createBookRequest struct {
	Isbn        string    `json:"isbn" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Author      string    `json:"author" binding:"required"`
	ImageUrl    string    `json:"image_url" binding:"required"`
	Genre       string    `json:"genre" binding:"required"`
	Quantity    int       `json:"quantity" binding:"required"`
	PublishedAt time.Time `json:"published_at" binding:"required"`
}

func (server *Server) createBook(c *gin.Context) {
	var req createBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// TODO: Add upload to s3

	arg := db.CreateBookParams{
		Isbn:  req.Isbn,
		Title: req.Title,
		Description: pgtype.Text{
			String: req.Description,
			Valid:  req.Description != "",
		},
		Author: req.Author,
		ImageUrl: pgtype.Text{
			String: req.ImageUrl,
			Valid:  req.ImageUrl != "",
		},
		Genre:       req.Genre,
		Quantity:    int32(req.Quantity),
		PublishedAt: req.PublishedAt,
	}

	book, err := server.db.CreateBook(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, book)
}
