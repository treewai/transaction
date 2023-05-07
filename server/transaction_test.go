package server

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/treewai/transaction/db"
)

var s *Server

func TestMain(m *testing.M) {
	s = NewServer(db.NewDatabase())
}

func TestTransactionHandler(t *testing.T) {
	t.Run("get all transactions success", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/transactions", nil)
		if err != nil {
			t.Error(err)
		}
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(s.TransactionsHandler)
		handler.ServeHTTP(resp, req)
		if status := resp.Code; status != http.StatusOK {
			t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
		}
	})

	t.Run("add transaction success", func(t *testing.T) {
		jsonStr := []byte(`{"id":1,"account":"ABC","type":"Withdrawal","amount":100,"date":"2023-05-08T01:17:43.7598293+07:00","user":"user1"}`)

		req, err := http.NewRequest(http.MethodPost, "/transactions", bytes.NewBuffer(jsonStr))
		if err != nil {
			t.Error(err)
		}
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(s.TransactionsHandler)
		handler.ServeHTTP(resp, req)
		fmt.Println(resp.Code)
		if status := resp.Code; status != http.StatusOK {
			t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
		}
		expected := []byte(`{"id":1,"account":"ABC","type":"Withdrawal","amount":100,"date":"2023-05-08T01:17:43.7598293+07:00","user":"user1"}`)
		if !bytes.Equal(resp.Body.Bytes(), expected) {
			t.Errorf("handler returned unexpected body: got %v want %v",
				resp.Body.Bytes(), expected)
		}
	})
}

func TestTransactionByIdHandler(t *testing.T) {
	t.Run("get transaction by ID success", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/transactions/1", nil)
		if err != nil {
			t.Error(err)
		}
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(s.TransactionByIdHandler)
		handler.ServeHTTP(resp, req)
		if status := resp.Code; status != http.StatusOK {
			t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
		}
		expected := []byte(`{"id":1,"account":"ABC","type":"Withdrawal","amount":100,"date":"2023-05-08T01:17:43.7598293+07:00","user":"user1"}`)
		if !bytes.Equal(resp.Body.Bytes(), expected) {
			t.Errorf("handler returned unexpected body: got %v want %v",
				resp.Body.Bytes(), expected)
		}
	})
}
