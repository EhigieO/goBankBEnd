package services

import (
	"fmt"
	"rovaBankProject/dtos"
	"rovaBankProject/models"
	"rovaBankProject/repos"
	"rovaBankProject/utils"
)

type AccountServiceImpl struct {
	repository repos.AccountDb
}

func (accountService AccountServiceImpl) CreateAccount(accountRequest dtos.UserInfo) (*models.Account, error) {
	//if accountRequest.InitialCredit > 0 {
	//	err := transactionService.repository.CreateTransaction()
	//	if err != nil {
	//		return nil, err
	//	}
	//}
	account := models.Account{
		AccountNumber:  utils.GenerateAccountNumber(),
		AccountBalance: accountRequest.InitialCredit,
		CustomerID:     accountRequest.CustomerID,
		AccountType:    "Current",
		Transactions:   []models.Transaction{},
	}
	if err := accountService.repository.CreateAccount(account); err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return &account, nil
	}
}

func (accountService AccountServiceImpl) GetAccount(accountNumber string) (*models.Account, error) {
	account, err := accountService.repository.GetAccount(accountNumber)
	if err != nil {
		return nil, err
	}
	return account, nil
}
