package services

import (
	"rovaBankProject/dtos"
	"rovaBankProject/repos"
	"testing"
)

func TestCreateTransaction(t *testing.T) {
	accountRequest := dtos.UserInfo{
		CustomerID:    "123456",
		InitialCredit: 100,
	}
	repository := &repos.DB{}
	accountService := AccountServiceImpl{
		repository: repository,
	}

	transactionService := TransactionServiceImpl{
		repository: repository,
	}

	account, _ := accountService.CreateAccount(accountRequest)

	amount := 230010.0

	transaction, err := transactionService.CreateTransaction("initial deposit", account.AccountNumber, amount)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if transaction == nil {
		t.Errorf("transaction not created")
	}
}
