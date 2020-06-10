package main

import(
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
	http.HandleFunc(root + "/movie", movieHandler)
	return http.ListenAndServe(fullAddr, nil)
}

func movieHandler(w http.ResponseWriter, r *http.Request) {

}