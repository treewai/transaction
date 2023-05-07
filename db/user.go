package db

import (
	"github.com/treewai/transaction/models"
)

func (db *database) IsValidUser(u *models.User) bool {
	for _, user := range db.users {
		if user.Username == u.Username && user.Password == u.Password {
			return true
		}
	}
	return false
}
