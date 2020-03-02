package handler

import "net/http"

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	GetDetails() Details
}

type Details struct {
	Name string
	Method string
	Endpoint string
}