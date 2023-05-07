package db

import "github.com/treewai/transaction/models"

type Database interface {
	IsValidUser(*models.User) bool

	GetAllTransactions() []models.Transaction
	GetTransaction() models.Transaction
	AddTransaction(*models.Transaction) error
	UpdateTransaction(*models.Transaction) error
}

type database struct {
	users        []models.User
	transactions []models.Transaction
}

func NewDatabase() Database {
	return &database{}
}
