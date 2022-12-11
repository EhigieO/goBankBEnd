package services

import (
	"errors"
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

func (accountService AccountServiceImpl) UpdateAccount(transaction models.Transaction) error {
	toAccount, err := accountService.GetAccountByAccNum(transaction.To)
	if err != nil {
		return err
	}
	fromAccount, err := accountService.GetAccountByAccNum(transaction.From)
	if err != nil {
		return err
	}
	if fromAccount.AccountBalance < transaction.Amount {
		return errors.New("insufficient balance")
	}

	toAccount.AccountBalance += transaction.Amount
	fromAccount.AccountBalance -= transaction.Amount

	toAccount.Transactions = append(toAccount.Transactions[:], transaction)
	fromAccount.Transactions = append(fromAccount.Transactions[:], transaction)

	accountService.repository.SaveAccount(*toAccount)
	accountService.repository.SaveAccount(*fromAccount)
	return nil
}

func (accountService AccountServiceImpl) GetAccount(customerID string) (*models.Account, error) {
	account, err := accountService.repository.GetAccount(customerID)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (accountService AccountServiceImpl) GetAccountByAccNum(accountNumber string) (*models.Account, error) {
	account, err := accountService.repository.GetAccountByAccNum(accountNumber)
	if err != nil {
		return nil, err
	}
	return account, nil
}
