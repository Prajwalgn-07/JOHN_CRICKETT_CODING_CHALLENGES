package main

import (
	"log"
	"src/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/tokenize", handlers.DataEncoder)
	router.POST("/detokenize", handlers.DataDecoder)
	router.POST("/token", handlers.CreateToken)
	// Run the server with HTTPS
	err := router.RunTLS(":8080", "certs/cert.pem", "certs/key.pem")
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
