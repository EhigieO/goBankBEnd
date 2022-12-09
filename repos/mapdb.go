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
		return fmt.Errorf("error saving account")
	}

}
