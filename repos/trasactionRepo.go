package repos

import (
	"errors"
	"rovaBankProject/models"
)

type TransactionDB interface {
	CreateTransaction(transaction models.Transaction) error
}

func (db *DB) CreateTransaction(transaction models.Transaction) error {
	if (*db)[transaction.ID] == nil {
		(*db)[transaction.ID] = transaction
		return nil
	} else {
		return errors.New("error: possible duplicate transaction")
	}
}
