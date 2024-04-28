package main

import (
	"log"
	"net/http"
)

type ApiServer struct {
	addr string
}

func NewApiServer(addr string) *ApiServer {
	return &ApiServer{
		addr: addr,
	}
}

func (s *ApiServer) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("teste"))
	})
	router.HandleFunc("GET /clients", GetClients)
	router.HandleFunc("GET /clients/{id}", GetClientById)
	router.HandleFunc("POST /clients", PostClient)

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Printf("Listening on: %s", s.addr)

	return server.ListenAndServe()
}
