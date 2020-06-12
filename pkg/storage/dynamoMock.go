package storage

import (
	"fmt"
	"github.com/Perezonance/movie-manager-service/pkg/models"
	"github.com/Perezonance/movie-manager-service/pkg/primatives"
	"github.com/Perezonance/movie-manager-service/pkg/primitives"
)

type (
	DynamoMock struct {
		users map[float64]models.User
		posts map[float64]models.Post
	}
)

func NewMockDynamo() (*DynamoMock, error){
	return &DynamoMock{}, nil
}

//TODO:Implement Business Logic

/////////////////////////////////// User Services ///////////////////////////////////

func (d *DynamoMock)GetUser(id float64) (models.User, error){
	fmt.Printf(logRoot + "Searching %v table for user with id:%v\n", userTable, id)

	var (
		user = 	models.User{}
		err  =	primitives.ErrUserNotFound
	)
	user = d.users[id]
	if(user != models.User{}) {
		err = nil
	}
	return user, err
}

func (d *DynamoMock)PostUser(user models.User) error {
	fmt.Printf(logRoot + "Inserting user into %v table:%v\n", userTable, user)
	d.users[user.Id] = user
	return nil
}
func (d *DynamoMock)DeleteUser(user models.User) (models.User, error) {
	fmt.Printf(logRoot + "Deleting user from %v table:%v\n", userTable, user)
	fmt.Println("UNIMPLEMENTED")
	return user, nil
}

/////////////////////////////////// Post Services ///////////////////////////////////

func (d *DynamoMock)GetPost(id float64) (models.Post, error){
	fmt.Printf(logRoot + "Searching %v table for post with id:%v\n", userTable, id)

	var (
		post = models.Post{}
		err = primatives.ErrPostNotFound
	)

	post = d.posts[id]
	if(post != models.Post{}) {
		err = nil
	}
	return post, err
}

func (d *DynamoMock)PostPost(user models.User) error {
	fmt.Printf(logRoot + "Inserting post into %v table:%v\n", postTable, user)
	fmt.Println("UNIMPLEMENTED")
	return nil
}
func (d *DynamoMock)DeletePost(user models.User) (models.User, error) {
	fmt.Printf(logRoot + "Deleting post from %v table:%v\n", postTable, user)
	fmt.Println("UNIMPLEMENTED")
	return user, nil
}