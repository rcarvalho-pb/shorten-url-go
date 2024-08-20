package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	data "github.com/rcarvalho-pb/link-shorter/internal/models"
)

func NewURL(w http.ResponseWriter, r *http.Request) {
	log.Println("Here")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	defer r.Body.Close()

	var req data.UrlRequest

	if err = json.Unmarshal(body, &req); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err)
		return
	}

	var url data.ShortURL

	newURL, err := url.NewShortURL(req.URL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"url": newURL})
}
