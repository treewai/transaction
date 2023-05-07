package models

import "time"

type Transaction struct {
	ID      int       `json:"id"`
	Account string    `json:"account"`
	Type    string    `json:"type"`
	Amount  float64   `json:"amount"`
	Date    time.Time `json:"date"`
	User    string    `json:"user"`
}
