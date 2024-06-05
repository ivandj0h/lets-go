package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

type MySQLStorage struct {
	db *sql.DB
}

func NewMySQLStorage(cfg mysql.Config) *MySQLStorage {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("MySQL has been Connected!")
	return &MySQLStorage{db: db}
}

func (s *MySQLStorage) Init() (*sql.DB, error) {

	// table init

	return s.db, nil
}
