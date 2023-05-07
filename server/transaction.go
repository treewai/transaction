package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/treewai/transaction/db"
	"github.com/treewai/transaction/models"
)

func (s *Server) TransactionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		transactions := s.db.GetAllTransactions()
		if err := json.NewEncoder(w).Encode(transactions); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return

	case http.MethodPost:
		var t models.Transaction
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if t.Type != "Deposit" && t.Type != "Withdrawal" {
			http.Error(w, "Type not support", http.StatusBadRequest)
			return
		}

		if t.Date.IsZero() {
			t.Date = time.Now()
		}

		tx, err := s.db.AddTransaction(&t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(tx); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "Not Found", http.StatusNotFound)
}

func (s *Server) TransactionByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.URL.Path[len("/transactions/"):])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		tx, err := s.db.GetTransaction(id)
		if err != nil {
			if err == db.ErrIDNotFound {
				http.NotFound(w, r)
				return
			}

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(tx); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return

	case http.MethodPut:
		var t models.Transaction
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		t.ID = id
		tx, err := s.db.UpdateTransaction(&t)
		if err != nil {
			if err == db.ErrIDNotFound {
				http.NotFound(w, r)
				return
			}

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(tx); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return

	case http.MethodDelete:
		tx, err := s.db.DeleteTransaction(id)
		if err != nil {
			if err == db.ErrIDNotFound {
				http.NotFound(w, r)
				return
			}

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(tx); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.NotFound(w, r)
}
