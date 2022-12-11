package utils

import (
	"log"
	"math/rand"
	"os"
	"time"
)

func GenerateAccountNumber() string {
	return idGen()
}
func GenerateTransactionId() string {
	return idGen()
}
func GenerateCustomerID() string {
	return idGen()
}

func idGen() string {
	var letter = []rune("0123456789")

	rand.Seed(time.Now().UnixNano())
	accountNumber := make([]rune, 10)
	for i := range accountNumber {
		accountNumber[i] = letter[rand.Intn(len(letter))]
	}
	return string(accountNumber)
}

var ApplicationLog = log.New(os.Stdout, "[account-service] ", log.LstdFlags)
