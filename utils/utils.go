package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"os"
	"rovaBankProject/response"
	"time"
)

func GenerateAccountNumber() string {
	var letter = []rune("0123456789")

	rand.Seed(time.Now().UnixNano())
	accountNumber := make([]rune, 10)
	for i := range accountNumber {
		accountNumber[i] = letter[rand.Intn(len(letter))]
	}
	return string(accountNumber)
}
func GenerateTransactionId() string {
	return GenerateCustomerID()
}

var customerCounter int64 = 1

func GenerateCustomerID() string {
	customerCounter += 1
	return fmt.Sprintf("%06d", customerCounter)
}

var ApplicationLog = log.New(os.Stdout, "[account-service] ", log.LstdFlags)

func GenerateJSONResponse(c *gin.Context, statusCode int, message string, data map[string]interface{}) {
	c.JSON(statusCode, response.APIResponse{
		Status:    statusCode,
		Message:   message,
		Timestamp: time.Now(),
		Data:      data,
	})
}

func GenerateInternalServerErrorResponse(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, response.APIResponse{
		Status:    http.StatusInternalServerError,
		Message:   message,
		Timestamp: time.Now(),
		Data:      gin.H{},
	})
}

func GenerateBadRequestResponse(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, response.APIResponse{
		Status:    http.StatusBadRequest,
		Message:   message,
		Timestamp: time.Now(),
		Data:      gin.H{},
	})
}
