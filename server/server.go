package server

import "github.com/treewai/transaction/db"

type Server struct {
	db db.Database
}

func NewServer(db db.Database) *Server {
	return &Server{
		db: db,
	}
}
