package main

import (
	"github.com/Perezonance/movie-manager-service/pkg/models"
	"github.com/Perezonance/movie-manager-service/pkg/server"
	"github.com/Perezonance/movie-manager-service/pkg/storage"

	"fmt"
	"net/http"
)

const (
	root = "/api/v1"
	port = "8080"
	addr = "0.0.0.0"
	fullAddr = addr + ":" + port
)

func main() {
	fmt.Println(start())
}

func start() error{
	db, err := storage.NewDynamo()
	if err != nil {
		fmt.Println(err)
	}

	s, err := server.NewServer(db)

	fmt.Println("Starting up server...")
	http.HandleFunc(root + "/user", server.UserHandler)
	http.HandleFunc(root + "/post", server.PostHandler)

	fmt.Printf("Listening on %v\n", fullAddr)
	return http.ListenAndServe(fullAddr, nil)
}

func PostUser(u models.User) {
	fmt.Printf("User with id %v and name %v has been added to DB\n", u.Id, u.Name)
}