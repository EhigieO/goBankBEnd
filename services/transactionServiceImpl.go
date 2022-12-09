package services

import (
	"rovaBankProject/models"
	"rovaBankProject/repos"
	"time"
)

type TransactionServiceImpl struct {
	repository repos.TransactionDB
}

func (transactionService *TransactionServiceImpl) CreateTransaction(fromAccount, toAccount string, amount float64) (*models.Transaction, error) {
	transaction := models.Transaction{
		Amount: amount,
		From:   fromAccount,
		To:     toAccount,
		Date:   time.Now().Local(),
	}
	err := transactionService.repository.CreateTransaction(transaction)

	if err != nil {
		return nil, err
	}
	return &transaction, nil
}
