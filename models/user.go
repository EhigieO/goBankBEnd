package models

type User struct {
	CustomerID string    `json:"customerID"`
	Name       string    `json:"name"`
	Surname    string    `json:"surname"`
	Accounts   []Account `json:"accounts"`
}
