package api

import (
	"net/http"
	"time"

	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

type GetMemberRequest struct {
	ID string `uri:"id"`
}

func (server *Server) getMember(c *gin.Context) {
	var req GetMemberRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	member, err := server.db.GetMember(c, uuid.MustParse(req.ID))
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, member)
}

func (server *Server) listMembers(c *gin.Context) {
	member, err := server.db.ListMembers(c)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, member)
}

type CreateMemberRequest struct {
	Email     string `json:"email" binding:"required,email"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"`
	Dob       int    `json:"dob" binding:"required"`
	Gender    string `json:"gender" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type CreateMemberResponse struct {
}

func (server *Server) createMember(c *gin.Context) {
	var req CreateMemberRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// TODO: Add validator for request fields
	// TODO: Refactor to achieve modularity. Single Responsibility Principle
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateMemberParams{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName: pgtype.Text{
			String: req.LastName,
			Valid:  req.LastName != "",
		},
		Dob: pgtype.Date{
			Time:  time.UnixMilli(int64(req.Dob)),
			Valid: true,
		},
		Gender:   db.Gender(req.Gender),
		Password: string(hashedPassword),
	}

	member, err := server.db.CreateMember(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, member)
}
