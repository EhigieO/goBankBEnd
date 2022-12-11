package services

import (
	"fmt"
	"rovaBankProject/dtos"
	"rovaBankProject/models"
	"rovaBankProject/repos"
	"rovaBankProject/utils"
)

type UserServiceImpl struct {
	repository repos.UserDB
}

func (userService UserServiceImpl) CreateUser(userRequest dtos.NewUser) (*models.User, error) {

	user := models.User{
		CustomerID: utils.GenerateCustomerID(),
		Name:       userRequest.Name,
		Surname:    userRequest.Surname,
		Accounts:   []models.Account{},
	}

	if err := userService.repository.CreateUser(user); err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return &user, nil
	}
}

func (userService UserServiceImpl) UpdateUser(customerID string, account models.Account) error {
	dbUser, err := userService.GetUser(customerID)
	if err != nil {
		return err
	}

	dbUser.Accounts = append(dbUser.Accounts[:], account)

	userService.repository.SaveUser(*dbUser)
	return nil
}

func (userService UserServiceImpl) GetUser(customerID string) (*models.User, error) {
	account, err := userService.repository.GetUser(customerID)
	if err != nil {
		return nil, err
	}
	return account, nil
}
