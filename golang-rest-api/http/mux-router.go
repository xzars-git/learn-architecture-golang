package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var (
	muxDisatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDisatcher.HandleFunc(uri, f).Methods("GET")
}
func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDisatcher.HandleFunc(uri, f).Methods("POST")
}
func (*muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server running on p[ort %v", port)
	http.ListenAndServe(port, muxDisatcher)
}
