package db

import (
	"errors"

	"github.com/treewai/transaction/models"
)

var (
	ErrIDNotFound = errors.New("ID Not Found")
)

func (db *database) GetAllTransactions() []models.Transaction {
	db.Lock()
	defer db.Unlock()

	transactions := []models.Transaction{}
	for _, v := range db.transactions {
		transactions = append(transactions, *v)
	}
	return transactions
}

func (db *database) GetTransaction(id int) (models.Transaction, error) {
	db.Lock()
	defer db.Unlock()

	t, ok := db.transactions[id]
	if !ok {
		return models.Transaction{}, ErrIDNotFound
	}
	return *t, nil
}

func (db *database) AddTransaction(t *models.Transaction) (models.Transaction, error) {
	db.Lock()
	defer db.Unlock()

	db.lastID++
	t.ID = db.lastID
	db.transactions[t.ID] = t
	return *t, nil
}

func (db *database) UpdateTransaction(t *models.Transaction) (models.Transaction, error) {
	db.Lock()
	defer db.Unlock()

	tx, ok := db.transactions[t.ID]
	if !ok {
		return models.Transaction{}, ErrIDNotFound
	}

	if t.Account != "" {
		tx.Account = t.Account
	}
	if t.Type != "" {
		tx.Type = t.Type
	}
	if t.Amount > 0.0 {
		tx.Amount = t.Amount
	}
	if !t.Date.IsZero() {
		tx.Date = t.Date
	}
	if t.User != "" {
		tx.User = t.User
	}

	return *t, nil
}

func (db *database) DeleteTransaction(id int) (models.Transaction, error) {
	db.Lock()
	defer db.Unlock()

	t, ok := db.transactions[id]
	if !ok {
		return models.Transaction{}, ErrIDNotFound
	}

	delete(db.transactions, id)
	return *t, nil
}
