package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type data struct {
	ID   string            `json:"id"`
	Data map[string]string `json:"data"`
}

func tokenize(c *gin.Context) {
	var request data
	if err := c.BindJSON(&request); err != nil {
		return
	}

	for key, value := range request.Data {
		request.Data[key] = value
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"status": "success", "message": "Data has been tokenized"})
}

func detokenize(c *gin.Context) {
	var request data

	if err := c.BindJSON(&request); err != nil {
		return
	}

	for key, value := range request.Data {
		request.Data[key] = value
	}

}

func main() {
	router := gin.Default()
	router.POST("/tokenize", tokenize)
	router.POST("/detokenize", detokenize)
	router.Run("localhost:8080")
}
