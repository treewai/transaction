package main

import (
	"log"
	"net/http"

	"github.com/treewai/transaction/db"
	"github.com/treewai/transaction/server"
)

func main() {
	s := server.NewServer(db.NewDatabase())

	http.HandleFunc("/auth", s.AuthHandler)
	http.HandleFunc("/transactions", s.TokenMiddleware(s.TransactionsHandler))
	http.HandleFunc("/transactions/", s.TokenMiddleware(s.TransactionByIdHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
