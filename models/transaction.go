package models

import "time"

type Transaction struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
	From   string  `json:"from"`
	To     string  `json:"to"`
	Date   time.Time
}
