package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `jdon:"password"`
}
