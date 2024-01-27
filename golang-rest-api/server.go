package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var router = mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Up and running...")
	})
	router.HandleFunc("/posts", getPost).Methods("GET")
	router.HandleFunc("/posts", addPost).Methods("POST")
	log.Println("Server started on port", port)
	log.Fatal(http.ListenAndServe(port, router))

}
