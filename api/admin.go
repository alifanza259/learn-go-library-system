package api

import (
	"fmt"
	"net/http"

	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
	"github.com/alifanza259/learn-go-library-system/util"
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

type LoginAdminRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginAdminResponse struct {
	Admin           adminResponse `json:"admin"`
	AccessToken     string        `json:"access_token"`
	AccessExpiredAt int           `json:"access_expired_at"`
}

type adminResponse struct {
	ID        uuid.UUID
	Email     string
	FirstName string
	LastName  pgtype.Text
	CreatedAt int64
}

func newAdminResponse(admin db.Admin) adminResponse {
	return adminResponse{
		ID:        admin.ID,
		Email:     admin.Email,
		FirstName: admin.FirstName,
		LastName:  admin.LastName,
		CreatedAt: admin.CreatedAt.Unix(),
	}
}

func (server *Server) loginAdmin(c *gin.Context) {
	var req LoginAdminRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	admin, err := server.db.GetAdminByEmail(c, req.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("please check your credentials")))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, admin.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("please check your credentials")))
		return
	}

	accessToken, accessExpiresAt, err := server.tokenMaker.CreateToken(admin.Email, server.config.AccessTokenDuration, "admin")
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := LoginAdminResponse{
		Admin:           newAdminResponse(admin),
		AccessToken:     accessToken,
		AccessExpiredAt: accessExpiresAt,
	}

	c.JSON(http.StatusOK, resp)
}

func (server *Server) listMembers(c *gin.Context) {
	member, err := server.db.ListMembers(c)
	if err != nil {

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, member)
}
