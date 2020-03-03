package home

import(
	"Allfather/rest/handler"
	"fmt"
	"net/http"
)

type Index struct{
	Message string
}

func (t Index) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, t.Message)
}

func (t Index) GetDetails() handler.Details {
	return handler.Details{
		Name:     "Index",
		Method:   "GET",
		Endpoint: "/",
	}
}