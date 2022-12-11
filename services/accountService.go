package services

import (
	"rovaBankProject/dtos"
	"rovaBankProject/models"
)

type AccountService interface {
	CreateAccount(accountRequest dtos.UserInfo) (*models.Account, error)
	GetAccount(accountNumber string) (*models.Account, error)
	UpdateAccount(transaction models.Transaction) error
	GetAccountByAccNum(accountNumber string) (*models.Account, error)
}
