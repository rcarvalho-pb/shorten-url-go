package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func routes() http.Handler {
	router := mux.NewRouter()
	configRouter(router)
	return router
}
