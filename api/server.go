package api

import (
	db "github.com/AnkitNayan83/backend-boilerplate-go/db/sqlc"
	"github.com/AnkitNayan83/backend-boilerplate-go/token"
	"github.com/AnkitNayan83/backend-boilerplate-go/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store      db.Store
	router     *gin.Engine
	config     util.Config
	tokenMaker token.Maker
}

func NewServer(store db.Store, config util.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.PasetoSymmetricKey)
	if err != nil {
		return nil, err
	}
	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: tokenMaker,
	}

	server.setupServerRoutes()

	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) setupServerRoutes() {
	router := gin.Default()

	v1Router := router.Group("/api/v1")

	// health check
	v1Router.GET("/health", server.healthCheck)
	// user routes
	v1Router.GET("/user/:id", server.getUserById)
	// auth routes

	v1Router.POST("/auth/register", server.createUser)
	v1Router.POST("/auth/login", server.credentialLogin)

	server.router = router
}
