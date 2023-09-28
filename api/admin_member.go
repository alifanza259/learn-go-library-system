package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) listMembers(c *gin.Context) {
	member, err := server.db.ListMembers(c)
	if err != nil {

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, member)
}
