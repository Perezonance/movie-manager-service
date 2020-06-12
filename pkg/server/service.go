package server

import (
	"encoding/json"
	"github.com/Perezonance/movie-manager-service/pkg/models"
	"github.com/Perezonance/movie-manager-service/pkg/storage"
	"net/http"
)

type (
	Server struct {
		db *storage.DynamoDB
	}
)

/////////////////////////////////// User Service Functions ///////////////////////////////////

func NewServer(db Persistence{}) (Server, error) {
	//TODO: Server instance id setup
	return Server{db:db}, nil
}

//TODO: need to allow both an array or a single json object as request payload
func (s *Server)getUser(w http.ResponseWriter, r *http.Request) {
	writeRes(200, "UNIMPLEMENTED", w)
}

func (s *Server)postUser(w http.ResponseWriter, r *http.Request) {
	var reqPayload models.RequestUserPayload
	err := json.NewDecoder(r.Body).Decode(&reqPayload)
	if err != nil {
		//TODO: ERROR HANDLING
	}

	for _, u := range reqPayload.Payload {
		go func(){
			err := s.db.PostUser(u)
			if err != nil {
				//TODO: ERROR HANDLING
			}
		}()
	}
	writeRes(http.StatusOK, http.StatusText(http.StatusOK), w)
}

func (s *Server)deleteUser(w http.ResponseWriter, r *http.Request) {
	writeRes(200, "UNIMPLEMENTED", w)
}

/////////////////////////////////// Post Service Functions ///////////////////////////////////

func (s *Server)getPost(w http.ResponseWriter, r *http.Request) {
	writeRes(200, "UNIMPLEMENTED", w)
}

func (s *Server)postPost(w http.ResponseWriter, r *http.Request) {
	var reqPayload models.RequestPostPayload
	err := json.NewDecoder(r.Body).Decode(&reqPayload)
	if err != nil {
		//TODO: ERROR HANDLING
	}

	for _, p := range reqPayload.Payload {
		go func() {
			err := s.db.PostPost(p)
			if err != nil {
				//TODO: ERROR HANDLING
			}
		}()
	}
	writeRes(http.StatusOK, http.StatusText(http.StatusOK), w)
}

func (s *Server)deletePost(w http.ResponseWriter, r *http.Request) {
	writeRes(200, "UNIMPLEMENTED", w)
}