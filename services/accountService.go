package services

import (
	"rovaBankProject/dtos"
	"rovaBankProject/models"
)

type AccountService interface {
	CreateAccount(accountRequest dtos.UserInfo) (*models.Account, error)
}