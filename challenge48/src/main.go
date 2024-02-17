package main

import (
	"src/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/tokenize", handlers.DataEncoder)
	router.POST("/detokenize", handlers.DataDecoder)
	router.POST("/token", handlers.CreateToken)
	router.Run(":8080")
}
