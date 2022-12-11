package services

import (
	"errors"
	"rovaBankProject/models"
	"rovaBankProject/repos"
	"rovaBankProject/utils"
	"time"
)

type TransactionServiceImpl struct {
	repository repos.TransactionDB
}

func (transactionService *TransactionServiceImpl) CreateTransaction(fromAccount, toAccount string, amount float64) (*models.Transaction, error) {
	if amount <= 0 {
		return nil, errors.New("amount too small")
	}
	if toAccount == "" {
		return nil, errors.New("account  number cannot be empty")
	}
	if fromAccount == "" {
		return nil, errors.New("account  number cannot be empty")
	}
	transaction := models.Transaction{
		ID:     utils.GenerateTransactionId(),
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
