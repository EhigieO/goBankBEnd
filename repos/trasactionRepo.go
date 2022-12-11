package repos

import (
	"errors"
	"rovaBankProject/models"
)

type TransactionDB interface {
	CreateTransaction(transaction models.Transaction) error
}

type TDB map[string]interface{}

func (db *TDB) CreateTransaction(transaction models.Transaction) error {
	if (*db)[transaction.ID] == nil {
		(*db)[transaction.ID] = transaction
		return nil
	} else {
		return errors.New("error: possible duplicate transaction")
	}
}
