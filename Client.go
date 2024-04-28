package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Client struct {
	ID       int      `json: "id"`
	Fullname string   `json: "fullname"`
	Contats  []Contat `json: "contats"`
}

type Contat struct {
	Type  string `json: "type"`
	Value string `json: "value"`
}

func GetClients(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(clients)
}

func GetClientById(w http.ResponseWriter, r *http.Request) {
	clientId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		panic("conversion error")
	}

	for _, client := range clients {
		if clientId == client.ID {
			json.NewEncoder(w).Encode(client)
			return
		}
	}
	json.NewEncoder(w).Encode(&Client{})
}

func PostClient(w http.ResponseWriter, r *http.Request) {
	var client Client
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		log.Fatal("Erro de decodificação de JSON")
		return
	}
	clients = append(clients, client)
	json.NewEncoder(w).Encode(clients)
}
