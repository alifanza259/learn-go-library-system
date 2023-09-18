package api

import (
	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
	"github.com/alifanza259/learn-go-library-system/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	db     db.Library
	config util.Config
	router *gin.Engine
}

func NewServer(db db.Library, config util.Config) (*Server, error) {
	server := &Server{db: db, config: config}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	r := gin.Default()

	r.GET("/admin/:id", server.getAdmin)
	r.GET("/admin", server.listAdmin)

	r.GET("/member/:id", server.getMember)
	r.GET("/members", server.listMembers)
	r.POST("/members", server.createMember)
	r.POST("/members/login", server.loginMember)
	server.router = r
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
