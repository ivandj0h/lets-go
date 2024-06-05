package main

type APIServer struct {
	addr  string
	store Store
}

func NewApiServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}
