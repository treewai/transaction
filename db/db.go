package db

import (
	"sync"

	"github.com/treewai/transaction/models"
)

type Database interface {
	IsValidUser(*models.User) bool

	GetAllTransactions() []models.Transaction
	GetTransaction(int) (models.Transaction, error)
	AddTransaction(*models.Transaction) error
	UpdateTransaction(*models.Transaction) error
	DeleteTransaction(int) error
}

type database struct {
	users        []models.User
	transactions map[int]*models.Transaction
	lastID       int
	sync.Mutex
}

func NewDatabase() Database {
	return &database{
		transactions: make(map[int]*models.Transaction),
	}
}
