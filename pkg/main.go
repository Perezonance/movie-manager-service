package main

import(
	"encoding/json"
	"github.com/Perezonance/movie-manager-service/pkg/models"

	"fmt"
	"io/ioutil"
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

	fmt.Printf("Listening on %v\n", fullAddr)
	return http.ListenAndServe(fullAddr, nil)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//TODO: Error Handling
		fmt.Println(err)
	}

	var payload = models.RequestUserPayload{}
	err = json.Unmarshal(body, &payload)
	if err != nil {
		//TODO: Error Handling
		fmt.Println(err)
	}

	for _, u := range payload.Payload {
		fmt.Println(u.Name)
		go PostUser(u)
	}
	fmt.Printf("Recieved request to /user endpoint\n")
	//fmt.Printf("Received request to /movie endpoint with the following payload: \n%v\n", string(body))
}

func postHandler(w http.ResponseWriter, r *http.Request) {

}

func PostUser(u models.User) {
	fmt.Printf("User with id %v and name %v has been added to DB\n", u.Id, u.Name)
}