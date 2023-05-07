package models

import "time"

type Transaction struct {
	ID      int       `json:"id,omitempty"`
	Account string    `json:"account"`
	Type    string    `json:"type"`
	Amount  float64   `json:"amount"`
	Date    time.Time `json:"date,omitempty"`
	User    string    `json:"user"`
}
