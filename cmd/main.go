package main

import (
	"github.com/rcarvalho-pb/link-shorter/internal/db"
	data "github.com/rcarvalho-pb/link-shorter/internal/models"
	"github.com/rcarvalho-pb/link-shorter/internal/server"
)

func main() {

	db := db.InitDB()

	models := data.New(db)

	go models.ShortURL.DeleteOldURLs()

	server.Serve()

}
