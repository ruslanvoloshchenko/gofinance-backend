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

	router.POST("/user", server.createUser)        // vamos por nossas rotas
	router.GET("/user/:username", server.getUser)  // vamos por nossas rotas
	router.GET("/user/id/:id", server.getUserById) // vamos por nossas rotas

	router.POST("/category", server.createCategory)       // vamos por nossas rotas
	router.GET("/category/id/:id", server.getCategory)    // vamos por nossas rotas
	router.GET("/category", server.getCategories)         // vamos por nossas rotas
	router.DELETE("/category/:id", server.deleteCategory) // vamos por nossas rotas
	router.PUT("/category/:id", server.updateCategory)    // vamos por nossas rotas

	router.POST("/account", server.createAccount)                            // vamos por nossas rotas
	router.GET("/account/id/:id", server.getAccount)                         // vamos por nossas rotas
	router.GET("/account", server.getAccounts)                               // vamos por nossas rotas
	router.GET("/account/graph/:user_id/:type", server.getAccountGraph)      // vamos por nossas rotas
	router.GET("/account/reports/:user_id/:type", server.getAccountsReports) // vamos por nossas rotas
	router.DELETE("/account/:id", server.deleteAccount)                      // vamos por nossas rotas
	router.PUT("/account/:id", server.updateAccount)                         // vamos por nossas rotas

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"api has error:": err.Error()}
}
