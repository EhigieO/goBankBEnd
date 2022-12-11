package services

import (
	"rovaBankProject/dtos"
	"rovaBankProject/models"
)

type UserService interface {
	CreateUser(userRequest dtos.NewUser) (*models.User, error)
	GetUser(customerID string) (*models.User, error)
	UpdateUser(customerID string, account models.Account) error
}
