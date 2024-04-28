package main

var clients []Client

func main() {
	// seed clients
	clients = append(clients, Client{ID: 1, Fullname: "Joaquim Fernandes", Contats: []Contat{{Type: "phone", Value: "(63) 99999-9999"}, {Type: "email", Value: "joaquim@teste.com"}}})
	clients = append(clients, Client{ID: 2, Fullname: "Larissa Menina", Contats: []Contat{{Type: "phone", Value: "(63) 99999-9999"}, {Type: "email", Value: "larissa@teste.com"}}})

	server := NewApiServer(":3000")
	server.Run()
}
