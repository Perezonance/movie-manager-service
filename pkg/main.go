package main

import(
	"encoding/json"
	"github.com/Perezonance/movie-manager-service/pkg/models"
	"time"

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
	fmt.Println("Starting up server..")
	http.HandleFunc(root + "/user", userHandler)
	http.HandleFunc(root + "/post", postHandler)
	http.HandleFunc(root + "/fruit", fruitHandler)

	fmt.Printf("Listening on %v\n", fullAddr)
	return http.ListenAndServe(fullAddr, nil)
}

func userHandler(w http.ResponseWriter, r *http.Request) {

	var reqPayload models.RequestUserPayload
	if err := json.NewDecoder(r.Body).Decode(&reqPayload); err != nil {
		//TODO: Error Handling
		fmt.Printf("Server error> %v\n", err)
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Error: %v", err)))
		return
	}

	//fmt.Printf("First user's name:%v\n", reqPayload.Payload[0].Name)

	for _, u := range reqPayload.Payload {
		go PostUser(u)
	}


	fmt.Printf("Recieved request to /user endpoint\n")
	time.Sleep(2 * time.Second)
	//fmt.Printf("Received request to /movie endpoint with the following payload: \n%v\n", string(body))
}

func postHandler(w http.ResponseWriter, r *http.Request) {

}

func fruitHandler(w http.ResponseWriter, r *http.Request) {
	var reqPayload models.TestFruitPayload
	if err := json.NewDecoder(r.Body).Decode(&reqPayload); err != nil {
		//TODO: Error Handling
		fmt.Printf("Server error> %v\n", err)
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Error: %v", err)))
		return
	}

	//fmt.Printf("First user's name:%v\n", reqPayload.Payload[0].Name)

	for _, u := range reqPayload {
		go PostUser(u)
	}


	fmt.Printf("Recieved request to /user endpoint\n")
	time.Sleep(2 * time.Second)
}

func PostUser(u models.User) {
	fmt.Printf("User with id %v and name %v has been added to DB\n", u.Id, u.Name)
}