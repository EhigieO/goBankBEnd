package services

import (
	"rovaBankProject/dtos"
	"rovaBankProject/repos"
	"testing"
)

func TestCreateAccountService(t *testing.T) {
	accountRequest := dtos.UserInfo{
		CustomerID:    "123456",
		InitialCredit: 100,
	}
	repository := &repos.DB{}
	accountService := AccountServiceImpl{
		repository: repository,
	}

	account, err := accountService.CreateAccount(accountRequest)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if account == nil {
		t.Errorf("Expected account to be non-nil, but got nil")
	}
	if account.AccountNumber == "" {
		t.Errorf("Expected account number to be non-empty, but got empty")
	}
	if account.AccountBalance != accountRequest.InitialCredit {
		t.Errorf("Expected account balance to be %v, but got %v", accountRequest.InitialCredit, account.AccountBalance)
	}
	if account.CustomerID != accountRequest.CustomerID {
		t.Errorf("Expected customer ID to be %v, but got %v", accountRequest.CustomerID, account.CustomerID)
	}
	if account.AccountType != "Current" {
		t.Errorf("Expected account type to be Current, but got %v", account.AccountType)
	}
	if len(account.Transactions) != 0 {
		t.Errorf("Expected transactions to be empty, but got %v", account.Transactions)
	}
}
