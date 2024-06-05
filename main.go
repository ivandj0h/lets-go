package main

import (
	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	cfg := mysql.Config{
		User:                 Envs.DBUser,
		Passwd:               Envs.DBPassword,
		Addr:                 Envs.DBAddress,
		DBName:               Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	sqlStorage := NewMYSQLStorage(cfg)

	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	store := NewStore(db)
	api := NewApiServer(":3000", store)
	api.Serve()
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Registering Services
	log.Println("Starting API Server at ", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
