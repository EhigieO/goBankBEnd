package services

import (
	"rovaBankProject/dtos"
	"rovaBankProject/models"
)

type AccountService interface {
	CreateAccount(accountRequest dtos.UserInfo) (*models.Account, error)
	GetAccount(accountNumber string) (*models.Account, error)
}
