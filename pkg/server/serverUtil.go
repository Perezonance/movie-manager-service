package server

import "net/http"

func writeRes(statusCode int, message string, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	res := []byte(message)
	w.Write(res)
}