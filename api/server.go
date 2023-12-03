package api

import (
	"fmt"

	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
	"github.com/alifanza259/learn-go-library-system/external"
	"github.com/alifanza259/learn-go-library-system/token"
	"github.com/alifanza259/learn-go-library-system/util"
	"github.com/alifanza259/learn-go-library-system/worker"
	"github.com/gin-gonic/gin"
)

type Server struct {
	db              db.Library
	config          util.Config
	router          *gin.Engine
	tokenMaker      token.Maker
	external        external.External
	taskDistributor worker.TaskDistributor
}

func NewServer(db db.Library, config util.Config, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.SymmetricKey, config.SymmetricKeyAdmin)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	external := external.NewAwsExternal(config.AwsSecretAccessKey, config.AwsAccessKeyId, config.AwsRegion)

	server := &Server{db: db, config: config, tokenMaker: tokenMaker, external: external, taskDistributor: taskDistributor}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	r := gin.Default()
	r.MaxMultipartMemory = server.config.MaxFileSize << 20

	v1Routes := r.Group("/v1")
	v1Routes.GET("/books", server.listBooks)
	v1Routes.GET("/books/:id", server.getBook)

	v1Routes.POST("/members", server.createMember)
	v1Routes.GET("/members/verify", server.verifyMember)
	v1Routes.POST("/members/login", server.loginMember)

	authRoutes := r.Group("/v1").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/member/:id", server.getMember)
	authRoutes.GET("/books/borrow_history", server.listBorrowRequests)
	authRoutes.POST("/books/borrow", server.borrowBooks)
	authRoutes.POST("/books/return", server.returnBooks)

	adminRoutes := r.Group("/admin")
	adminRoutes.POST("/admin/login", server.loginAdmin)

	adminAuthRoutes := r.Group("/admin").Use(adminAuthMiddleware(server.tokenMaker))
	adminAuthRoutes.GET("/admin/:id", server.getAdmin)
	adminAuthRoutes.GET("/admin", server.listAdmin)
	adminAuthRoutes.GET("/members", server.listMembers)
	adminAuthRoutes.POST("/books", server.createBook)
	adminAuthRoutes.PATCH("/book/:id", server.updateBook)
	adminAuthRoutes.DELETE("/book/:id", server.deleteBook)
	adminAuthRoutes.PATCH("/books/process_request", server.processBorrowReq)

	server.router = r
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
