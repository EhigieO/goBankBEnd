package repos

import (
	"fmt"
	"rovaBankProject/models"
)

type DB map[string]interface{}

func (db *DB) CreateAccount(account models.Account) error {
	if (*db)[account.CustomerID] == nil {
		(*db)[account.CustomerID] = account
		return nil
	} else {
		//utils.ApplicationLog.Printf("Error Saving account %v\n", error)
		return fmt.Errorf("error: error saving account")
	}
}

func (db *DB) GetAccount(customerID string) (*models.Account, error) {
	if val, ok := (*db)[customerID]; ok {
		account, ok := val.(models.Account)
		if !ok {
			return nil, fmt.Errorf("invalid account type")
		}
		return &account, nil
	} else {
		return nil, fmt.Errorf("account number does not exist")
	}
}
