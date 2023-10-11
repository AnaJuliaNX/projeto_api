package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var (
	muxRota = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}

}
func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxRota.HandleFunc(uri, f).Methods("GET")
}
func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxRota.HandleFunc(uri, f).Methods("POST")
}
func (*muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP serve executando na porta %v", port)
	http.ListenAndServe(port, muxRota)
}
