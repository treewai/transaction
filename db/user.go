package db

import (
	"github.com/treewai/transaction/models"
)

func getDefaultUsers() []models.User {
	return []models.User{
		{
			ID:       1,
			Username: "user1",
			Password: "password1",
		},
		{
			ID:       2,
			Username: "user2",
			Password: "password2",
		},
		{
			ID:       3,
			Username: "user3",
			Password: "password3",
		},
	}
}

func (db *database) IsValidUser(u *models.User) bool {
	for _, user := range db.users {
		if user.Username == u.Username && user.Password == u.Password {
			return true
		}
	}
	return false
}
