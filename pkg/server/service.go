package server

import (
	"github.com/Perezonance/movie-manager-service/pkg/storage"
)



type (
	Server struct {
		db storage.DynamoDB
	}
)

func NewServer(db storage.DynamoDB) (Server, error) {
	//TODO: Server instance id setup
	return Server{db:db}, nil
}