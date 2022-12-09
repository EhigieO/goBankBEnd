package services

import "rovaBankProject/models"

type TransactionService interface {
	CreateTransaction(fromAccount, toAccount string, amount float64) (*models.Transaction, error)
}
