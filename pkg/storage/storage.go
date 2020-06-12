package storage

import "github.com/Perezonance/movie-manager-service/pkg/models"

type persistence interface {
	GetUser(id string) (models.User, error)
	PostUser(user models.User) error
	DeleteUser(user models.User) error
	GetPost(id string) (models.Post, error)
	PostPost(user models.Post) error
	DeletePost(user models.Post) error
}
