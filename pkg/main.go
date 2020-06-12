package main

import (
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
	var s = server.Server{}
	if mock {
		db, err := storage.NewMockDynamo()
		if err != nil {
			//TODO: ERROR HANDLING
		}
		s, err := server.NewServer(db)
	} else {
		db, err := storage.NewDynamo()
		if err != nil {
			//TODO: ERROR HANDLING
		}
		s, err = server.NewServer(db)
	}

	fmt.Println("Starting up server...")
	http.HandleFunc(root + "/user", s.UserHandler)
	http.HandleFunc(root + "/post", s.PostHandler)

	fmt.Printf("Listening on %v\n", fullAddr)
	return http.ListenAndServe(fullAddr, nil)
}