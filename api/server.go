package api

import (
	db "github.com/AnkitNayan83/backend-boilerplate-go/db/sqlc"
	"github.com/AnkitNayan83/backend-boilerplate-go/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
	config util.Config
}

func NewServer(store db.Store, config util.Config) (*Server, error) {
	server := &Server{
		store:  store,
		config: config,
	}

	server.setupServerRoutes()

	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) setupServerRoutes() {
	router := gin.Default()

	// health check
	router.GET("/health", server.healthCheck)

	// users routes
	router.POST("/users", server.createUser)

	server.router = router
}
