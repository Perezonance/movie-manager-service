package server

import "net/http"

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUser(w, r)
	case http.MethodPut:
		postUser(w, r)
	case http.MethodPost:
		postUser(w, r)
	case http.MethodDelete:
		deleteUser(w, r)
	default:
		writeRes(http.StatusNotFound, http.StatusText(http.StatusNotFound), w)
	}
	return
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getPost(w, r)
	case http.MethodPut:
		postPost(w, r)
	case http.MethodPost:
		postPost(w, r)
	case http.MethodDelete:
	default:
		writeRes(http.StatusNotFound, http.StatusText(http.StatusNotFound), w)
	}
	return
}

func writeRes(statusCode int, message string, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	res := []byte(message)
	w.Write(res)
}


//TODO: redundant code in both handlers
//func userHandler(w http.ResponseWriter, r *http.Request) {
//	var reqPayload models.RequestUserPayload
//	if err := json.NewDecoder(r.Body).Decode(&reqPayload); err != nil {
//		//TODO: Error Handling
//		fmt.Printf("Server error> %v\n", err)
//		w.WriteHeader(500)
//		w.Write([]byte(fmt.Sprintf("Error: %v", err)))
//		return
//	}
//
//	fmt.Printf("Number of Users:%v\n", len(reqPayload.Payload))
//
//	for _, u := range reqPayload.Payload {
//		go PostUser(u)
//	}
//	fmt.Printf("Recieved request to /user endpoint\n")
//}
//
//func postHandler(w http.ResponseWriter, r *http.Request) {
//	var reqPayload models.RequestPostPayload
//	if err := json.NewDecoder(r.Body).Decode(&reqPayload); err != nil {
//		//TODO: Error Handling
//		fmt.Printf("Server error> %v\n", err)
//		w.WriteHeader(500)
//		w.Write([]byte(fmt.Sprintf("Error: %v", err)))
//		return
//	}
//
//	fmt.Printf("Number of Posts:%v\n", len(reqPayload.Payload))
//
//	for _, u := range reqPayload.Payload {
//		go PostUser(u)
//	}
//
//	fmt.Printf("Recieved request to /user endpoint\n")
//}