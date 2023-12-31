package api

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
	"github.com/alifanza259/learn-go-library-system/token"
	"github.com/alifanza259/learn-go-library-system/util"
	"github.com/alifanza259/learn-go-library-system/worker"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type GetMemberRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
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

	accessTokenPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	if member.Email != accessTokenPayload.Email {
		c.JSON(http.StatusUnauthorized, errorResponse(errors.New("unauthorized access")))
		return
	}

	c.JSON(http.StatusOK, member)
}

type CreateMemberRequest struct {
	Email     string `json:"email" binding:"required,email"`
	FirstName string `json:"first_name" binding:"required,max=30"`
	LastName  string `json:"last_name" binding:"max=30"`
	Dob       int    `json:"dob" binding:"required"`
	Gender    string `json:"gender" binding:"required"`
	Password  string `json:"password" binding:"required,min=6"`
}

func (server *Server) createMember(c *gin.Context) {
	var req CreateMemberRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	violation := validateCreateMemberRequest(req)
	if violation != nil {
		c.JSON(http.StatusBadRequest, errorResponse(violation))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
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

	member, err := server.db.CreateMemberTx(c, db.CreateMemberTxParams{
		CreateMemberParams: arg,
		AfterCreate: func(member db.Member) error {
			return server.taskDistributor.DistributeTaskSendVerifyEmail(c, &worker.PayloadSendVerifyEmail{
				Username: member.FirstName,
				UUID:     member.ID.String(),
			})
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, member)
}

func validateCreateMemberRequest(req CreateMemberRequest) error {
	if req.Dob > int(time.Now().UnixMilli()) {
		return errors.New("date of birth cannot be in the future")
	}
	// And more
	return nil
}

type LoginMemberRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginMemberResponse struct {
	Member          db.Member `json:"member"`
	AccessToken     string    `json:"access_token"`
	AccessExpiredAt int       `json:"access_expired_at"`
}

func (server *Server) loginMember(c *gin.Context) {
	var req LoginMemberRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	member, err := server.db.GetMemberByEmail(c, req.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("please check your credentials")))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, member.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("please check your credentials")))
		return
	}

	accessToken, accessExpiresAt, err := server.tokenMaker.CreateToken(member.Email, member.ID.String(), server.config.AccessTokenDuration, "member")
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := LoginMemberResponse{
		Member:          member,
		AccessToken:     accessToken,
		AccessExpiredAt: accessExpiresAt,
	}

	c.JSON(http.StatusOK, resp)
}

type VerifyMemberRequest struct {
	Token string `form:"token" binding:"required"`
}

func (server *Server) verifyMember(c *gin.Context) {
	var req VerifyMemberRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	emailVerif, err := server.db.GetEmailVerification(c, req.Token)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusBadRequest, errorResponse(errors.New("invalid token")))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if !emailVerif.IsUsed {
		err := server.db.VerifyEmailTx(c, db.VerifyEmailTxParams{
			EmailVerif: emailVerif,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "email verified"})
		return
	}

	c.JSON(http.StatusBadRequest, errors.New("token already used"))

}
