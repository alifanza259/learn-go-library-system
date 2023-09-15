package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type pingRequestParam struct {
	Q string `form:"q"`
}

func pingEndpoint(c *gin.Context) {
	var req pingRequestParam

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"q":       req.Q,
	})
}
