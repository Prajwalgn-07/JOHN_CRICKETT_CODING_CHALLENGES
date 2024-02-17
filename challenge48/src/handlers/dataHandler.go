package handlers

import (
	"net/http"
	"src/models"
	"src/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func DataEncoder(c *gin.Context) {
	if c.ContentType() != "application/json" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content-Type must be application/json"})
		return
	}

	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing token"})
		return
	}

	err := VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var request models.Data
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Missing required fields": err.Error()})
		return
	}

	var response models.Data
	response.ID = request.ID
	response.Data = make(map[string]string)
	redisClient := getRedisClient()
	defer redisClient.Close()

	for key, value := range request.Data {
		encrypted, _ := utils.Encrypt(value)
		setRedisData(redisClient, encrypted, key)
		response.Data[key] = encrypted
	}
	c.JSON(http.StatusOK, response)
}

func DataDecoder(c *gin.Context) {
	if c.ContentType() != "application/json" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content-Type must be application/json"})
		return
	}

	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing token"})
		return
	}

	err := VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var request models.Data
	var response models.ResponseData

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Missing required fields": err.Error()})
		return
	}

	response.ID = request.ID
	response.Data = make(map[string]models.DataField)
	redisClient := getRedisClient()
	defer redisClient.Close()

	for key, value := range request.Data {
		var decryptedRedisValue = getRedisData(redisClient, value)
		if decryptedRedisValue == "" {
			decryptedValue, _ := utils.Decrypt(value)
			if decryptedValue != "" {
				response.Data[key] = models.DataField{Found: true, Value: decryptedValue}
			} else if decryptedValue == "" {
				response.Data[key] = models.DataField{Found: false, Value: ""}
			}
		} else {
			response.Data[key] = models.DataField{Found: true, Value: decryptedRedisValue}
		}
	}
	c.JSON(http.StatusOK, response)
}
