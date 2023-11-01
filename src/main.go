package main

import (
	"net/http"
	"utils/handlers"
)

func main() {
	http.HandleFunc("/v1/upload", handlers.Upload)
	http.HandleFunc("/v1/all", handlers.GetAll)
	http.ListenAndServe("localhost:8000", nil)
}
