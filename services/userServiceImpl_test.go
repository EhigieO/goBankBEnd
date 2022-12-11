package services

import (
	"rovaBankProject/dtos"
	"rovaBankProject/repos"
	"testing"
)

func TestCreateUserService(t *testing.T) {
	userRequest := dtos.NewUser{
		Name:    "Ehi",
		Surname: "Ikp",
	}
	repository := &repos.UDB{}
	userService := UserServiceImpl{
		repository: repository,
	}

	user, err := userService.CreateUser(userRequest)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if user == nil {
		t.Errorf("Expected user to be non-nil, but got nil")
	}
	if user.CustomerID == "" {
		t.Errorf("Expected account number to be non-empty, but got empty")
	}
	if user.Name != userRequest.Name {
		t.Errorf("Expected saved user name to be %v, but got %v", user.Name, userRequest.Name)
	}

}
func TestGetUser(t *testing.T) {
	userRequest := dtos.NewUser{
		Name:    "Stanley",
		Surname: "Danzi",
	}
	repository := &repos.UDB{}
	userService := UserServiceImpl{
		repository: repository,
	}

	createdUSer, _ := userService.CreateUser(userRequest)

	result, err := userService.GetUser(createdUSer.CustomerID)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result == nil {
		t.Errorf("user not found")
	}
	if result.Name != createdUSer.Name {
		t.Errorf("names are not the same")
	}
}

func TestUpdateUser(t *testing.T) {
	userRequest := dtos.NewUser{
		Name:    "Stanley",
		Surname: "Danzi",
	}

	repository := &repos.UDB{}
	userService := UserServiceImpl{
		repository: repository,
	}

	createdUSer, _ := userService.CreateUser(userRequest)

	accountRequest := dtos.UserInfo{
		CustomerID:    createdUSer.CustomerID,
		InitialCredit: 100,
	}
	aRepository := &repos.DB{}
	accountService := AccountServiceImpl{
		repository: aRepository,
	}

	account, _ := accountService.CreateAccount(accountRequest)

	err := userService.UpdateUser(accountRequest.CustomerID, *account)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	result, _ := userService.GetUser(createdUSer.CustomerID)

	if len(result.Accounts) < 1 {
		t.Errorf("error: user not update")
	}
}
