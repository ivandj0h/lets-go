package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

type MYSQLStorage struct {
	db *sql.DB
}

func NewMYSQLStorage(cfg mysql.Config) *MYSQLStorage {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("MySQL has been Connected!")
	return &MYSQLStorage{db: db}
}

func (s *MYSQLStorage) Init() (*sql.DB, error) {

	// table init

	return s.db, nil
}
