package router

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

type chiRouter struct {
}

var (
	chiDispatcher = chi.NewRouter()
)

func (*chiRouter) GET(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	chiDispatcher.Get(uri, f)
}

func (*chiRouter) POST(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	chiDispatcher.Post(uri, f)
}

func (*chiRouter) SERVE(port string) {
	fmt.Printf("chi http server running on port : %v ", port)
	http.ListenAndServe(port, chiDispatcher)
}

func NewChiRouter() Router {
	return &chiRouter{}
}
