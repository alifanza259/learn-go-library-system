package api

import (
	"errors"
	"net/http"

	"github.com/alifanza259/learn-go-library-system/token"
	"github.com/alifanza259/learn-go-library-system/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (server *Server) listMembers(c *gin.Context) {
	accessPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	admin, err := server.db.GetAdmin(c, uuid.MustParse(accessPayload.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if authorized := util.IsAdminAuthorized(admin, adminPermissionSuper); !authorized {
		c.JSON(http.StatusForbidden, errorResponse(errors.New("admin is not authorized to list members")))
		return
	}

	member, err := server.db.ListMembers(c)
	if err != nil {

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, member)
}
