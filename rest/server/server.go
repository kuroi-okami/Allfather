package server

import (
	"Allfather/rest/handler"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func New() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	return router
}

func AddEndpoint(router *mux.Router, handler handler.Handler) {
	router.
		Methods(handler.GetDetails().Method).
		Path(handler.GetDetails().Endpoint).
		Name(handler.GetDetails().Name).
		Handler(handler)
}

func Serve(router *mux.Router, port string) {
	truePort := []string {
		":",
		port,
	}
	http.ListenAndServe(strings.Join(truePort, ""), router)
}
