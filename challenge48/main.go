package main

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
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

func encrypt(plaintext string) (string, error) {
	key := "my32digitkey12345678901234567890"
	iv := "my16digitIvKey12"

	var plainTextBlock []byte
	length := len(plaintext)

	if length%16 != 0 {
		extendBlock := 16 - (length % 16)
		plainTextBlock = make([]byte, length+extendBlock)
		copy(plainTextBlock[length:], bytes.Repeat([]byte{uint8(extendBlock)}, extendBlock))
	} else {
		plainTextBlock = make([]byte, length)
	}

	copy(plainTextBlock, plaintext)

	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(plainTextBlock))
	mode := cipher.NewCBCEncrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, plainTextBlock)

	str := base64.StdEncoding.EncodeToString(ciphertext)
	return str, nil
}

func decrypt(encrypted string) (string, error) {
	key := "my32digitkey12345678901234567890"
	iv := "my16digitIvKey12"

	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)

	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return "", err
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return "", fmt.Errorf("block size cant be zero")
	}

	mode := cipher.NewCBCDecrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, ciphertext)

	// Remove padding.
	padding := int(ciphertext[len(ciphertext)-1])
	if padding < 1 || padding > aes.BlockSize {
		return "", errors.New("padding error")
	}
	ciphertext = ciphertext[:len(ciphertext)-padding]

	return string(ciphertext), nil
}

// PKCS5UnPadding  pads a certain blob of data with necessary data to be used in AES block cipher
func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])

	return src[:(length - unpadding)]
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

func getRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}

func setRedisData(client *redis.Client, key string, value string) {
	ctx := context.Background()
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func getRedisData(client *redis.Client, key string) string {
	ctx := context.Background()
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return val
}

func main() {
	router := gin.Default()
	encrypted, _ := encrypt("test")
	decrypted, _ := decrypt(encrypted)

	fmt.Println(encrypted)
	fmt.Println(decrypted)
	router.POST("/tokenize", tokenize)
	router.POST("/detokenize", detokenize)
	router.Run(":8080")
}
