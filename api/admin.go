package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgtype"
)

type GetAdminRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type GetAdminResponse struct {
	ID         string      `json:"id" binding:"uuid"`
	Email      string      `json:"email"`
	FirstName  string      `json:"first_name"`
	LastName   pgtype.Text `json:"last_name"`
	Permission string      `json:"permission"`
	CreatedAt  string      `json:"created_at"`
	UpdatedAt  string      `json:"updated_at"`
}

func (server *Server) getAdmin(c *gin.Context) {
	var req GetAdminRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	admin, err := server.db.GetAdmin(c, uuid.MustParse(req.ID))
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := GetAdminResponse{
		ID:         admin.ID.String(),
		Email:      admin.Email,
		FirstName:  admin.FirstName,
		LastName:   admin.LastName,
		Permission: admin.Permission,
		CreatedAt:  admin.CreatedAt.String(),
		UpdatedAt:  admin.UpdatedAt.String(),
	}

	c.JSON(http.StatusOK, resp)
}

func (server *Server) listAdmin(c *gin.Context) {
	admin, err := server.db.ListAdmin(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, admin)
}
