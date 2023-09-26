package api

import (
	"errors"
	"net/http"
	"time"

	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
	external "github.com/alifanza259/learn-go-library-system/external/aws"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

	location, err := external.UploadToS3(server.config, file)
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
