package main

import "database/sql"

type Store interface {
	// CreateUser Users
	CreateUser() error
}

type Storage struct {
	db *sql.DB
}

// NewStore Constructor
func NewStore(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) CreateUser() error {
	return nil
}
