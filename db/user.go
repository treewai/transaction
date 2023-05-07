package db

import (
	"github.com/treewai/transaction/models"
)

func getUsers() []models.User {
	return []models.User{
		{1, "user1", "password1"},
		{2, "user2", "password2"},
		{3, "user3", "password3"},
		{4, "user4", "password4"},
		{5, "user5", "password5"},
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
