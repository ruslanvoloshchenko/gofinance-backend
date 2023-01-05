package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/wil-ckaew/gofinance-backend/db/sqlc"
)

type Server struct {
	store  *db.SQLStore
	router *gin.Engine
}

func NewServer(store *db.SQLStore) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/user", server.createUser)       // vamos por nossas rotas
	router.GET("/user/:username", server.getUser) // vamos por nossas rotas
	router.GET("/user/:id", server.getUserById)   // vamos por nossas rotas

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"api has error:": err.Error()}
}
