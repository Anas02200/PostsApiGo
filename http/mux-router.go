package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type muxRouter struct {
}

var (
	muxDispatcher = mux.NewRouter()
)

func (*muxRouter) GET(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc("/posts", f).Methods("POST")
}

func (*muxRouter) SERVE(port string) {
	fmt.Printf("mux http server running on port : %v ", port)
	http.ListenAndServe(port, muxDispatcher)
}

func NewMuxRouter() Router {
	return &muxRouter{}
}
