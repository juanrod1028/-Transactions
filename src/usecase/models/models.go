package models

import (
	"time"
)

type Transactions struct {
	Id           int       `json:"id"`
	Date         string    `json:"date"`
	Transactions string    `json:"transactions"`
	CreatedAt    time.Time `json:"createdAt"`
}

func NewTransaction(id int, date string, transactions string) *Transactions {
	return &Transactions{
		Id:           id,
		Date:         date,
		Transactions: transactions,
		CreatedAt:    time.Now().UTC(),
	}
}

type User struct {
	Identification string
	Email          string
	Transactions   []Transactions
}

func NewUser(identification string, email string, transactions []Transactions) *User {
	return &User{
		Identification: identification,
		Email:          email,
		Transactions:   transactions,
	}
}
