package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Account struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	AccNumber  int64     `json:"acc_number"`
	AccBalance int64     `json:"acc_balance"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewAccount(firstName string, lastName string) *Account {
	return &Account{
		// ID:        rand.Intn(10000),
		FirstName: firstName,
		LastName:  lastName,
		AccNumber: int64(rand.Intn(100000)),
		CreatedAt: time.Now().UTC(),
	}
}
