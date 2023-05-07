package db

import "github.com/treewai/transaction/models"

func (db *database) GetAllTransactions() []models.Transaction {
	// TODO
	return db.transactions
}

func (db *database) GetTransaction() models.Transaction {
	// TODO
	return models.Transaction{}
}

func (db *database) AddTransaction(*models.Transaction) error {
	// TODO
	return nil
}

func (db *database) UpdateTransaction(*models.Transaction) error {
	// TODO
	return nil
}

func (db *database) DeleteTransaction(id int) error {
	// TODO
	return nil
}
