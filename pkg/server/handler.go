package server

import (
	"encoding/json"
	"fmt"
	"github.com/Perezonance/movie-manager-service/pkg/models"
	"net/http"
)

//TODO: redundant code in both handlers
func userHandler(w http.ResponseWriter, r *http.Request) {
	var reqPayload models.RequestUserPayload
	if err := json.NewDecoder(r.Body).Decode(&reqPayload); err != nil {
		//TODO: Error Handling
		fmt.Printf("Server error> %v\n", err)
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Error: %v", err)))
		return
	}

	fmt.Printf("Number of Users:%v\n", len(reqPayload.Payload))

	for _, u := range reqPayload.Payload {
		go PostUser(u)
	}
	fmt.Printf("Recieved request to /user endpoint\n")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	var reqPayload models.RequestPostPayload
	if err := json.NewDecoder(r.Body).Decode(&reqPayload); err != nil {
		//TODO: Error Handling
		fmt.Printf("Server error> %v\n", err)
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Error: %v", err)))
		return
	}

	fmt.Printf("Number of Posts:%v\n", len(reqPayload.Payload))

	for _, u := range reqPayload.Payload {
		go PostUser(u)
	}

	fmt.Printf("Recieved request to /user endpoint\n")
}