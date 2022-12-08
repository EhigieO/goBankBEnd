package models

type Account struct {
	AccountNumber  string        `json:"accountNumber"`
	CustomerID     string        `json:"customerID"`
	AccountType    string        `json:"accountType"`
	AccountBalance float64       `json:"accountBalance"`
	Transactions   []Transaction `json:"transactions"`
}
