package repos

import (
	"errors"
	"rovaBankProject/models"
)

type AccountDb interface {
	CreateAccount(account models.Account) error
	GetAccount(customerID string) (*models.Account, error)
}

type DB map[string]interface{}

func (db *DB) CreateAccount(account models.Account) error {
	if (*db)[account.CustomerID] == nil {
		(*db)[account.CustomerID] = account
		return nil
	} else {
		return errors.New("error: error saving account")
	}
}

func (db *DB) GetAccount(customerID string) (*models.Account, error) {
	if val, ok := (*db)[customerID]; ok {
		account, ok := val.(models.Account)
		if !ok {
			return nil, errors.New("invalid account type")
		}
		return &account, nil
	} else {
		return nil, errors.New("account number does not exist")
	}
}
