package api

import (
	db "simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
)

// server serves http requests for our banking service

type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount) // http://localhost:8081/accounts
	router.GET("/accounts/:id", server.getAccount) // http://localhost:8081/accounts/105
	router.GET("/accounts", server.listAccount)    // http://localhost:8081/accounts/?page_id=1&page_size=5

	server.router = router
	return server
}

// Start runs the server on given address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
