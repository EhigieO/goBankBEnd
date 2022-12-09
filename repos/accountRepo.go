package repos

import "rovaBankProject/models"

type Repos interface {
	CreateAccount(account models.Account) error
	GetAccount(accountNumber string) (*models.Account, error)
}
