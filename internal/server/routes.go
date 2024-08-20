package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
}

var apiRoutes = []Route{
	{
		URI:      "/shorten-url",
		Method:   http.MethodPost,
		Function: NewURL,
	},
}

func configRouter(router *mux.Router) {

	apiRouter := router.PathPrefix("/api").Subrouter()
	urlApi := apiRouter.PathPrefix("/shorter").Subrouter()
	for _, r := range apiRoutes {
		urlApi.HandleFunc(r.URI, r.Function).Methods(r.Method)
	}
}
