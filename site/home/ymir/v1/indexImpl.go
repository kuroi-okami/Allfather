package ymir

import (
	"Allfather/rest/handler"
	"encoding/json"
	"net/http"
)

var Index handler.Handler = index{
	Message: indexData{
		Name:    "Ymir",
		Version: "V1",
		Methods: []string{
			"nil",
		},
	},
}

func (t index) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json := json.NewEncoder(w)

	if err := json.Encode(t); err != nil {
		panic(err)
	}
}

func (t index) GetDetails() handler.Details {
	return handler.Details{
		Name:     "Ymir/v1/Index",
		Method:   "GET",
		Endpoint: "/ymir/v1/",
	}
}
