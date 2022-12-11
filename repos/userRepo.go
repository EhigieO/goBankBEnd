package repos

import (
	"errors"
	"rovaBankProject/models"
)

type UserDB interface {
	CreateUser(user models.User) error
	SaveUser(user models.User)
	GetUser(customerID string) (*models.User, error)
}

type UDB map[string]interface{}

func (db *UDB) SaveUser(user models.User) {
	(*db)[user.CustomerID] = user
}

func (db *UDB) CreateUser(user models.User) error {
	if (*db)[user.CustomerID] == nil {
		(*db)[user.CustomerID] = user
		return nil
	} else {
		return errors.New("error: error saving account")
	}
}

func (db *UDB) GetUser(customerID string) (*models.User, error) {
	if val, ok := (*db)[customerID]; ok {
		user, ok := val.(models.User)
		if !ok {
			return nil, errors.New("invalid account type")
		}
		return &user, nil
	} else {
		return nil, errors.New("account number does not exist")
	}
}
