package api

import (
	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
	"github.com/alifanza259/learn-go-library-system/token"
	"github.com/alifanza259/learn-go-library-system/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	db         db.Library
	config     util.Config
	router     *gin.Engine
	tokenMaker token.Maker
}

func NewServer(db db.Library, config util.Config) (*Server, error) {
	tokenMaker := token.NewJWTMaker(config.SecretKey)
	server := &Server{db: db, config: config, tokenMaker: tokenMaker}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	r := gin.Default()

	r.GET("/admin/:id", server.getAdmin)
	r.GET("/admin", server.listAdmin)

	r.POST("/members", server.createMember)
	r.POST("/members/login", server.loginMember)

	authRoutes := r.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/member/:id", server.getMember)
	authRoutes.GET("/members", server.listMembers)

	server.router = r
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
