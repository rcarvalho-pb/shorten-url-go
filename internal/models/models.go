package data

import (
	"database/sql"
	"time"
)

var db *sql.DB
var dbTimeout = 10 * time.Second

type Model struct {
	ShortURL ShortURL
}

func New(dbPool *sql.DB) *Model {
	db = dbPool
	return &Model{
		ShortURL: ShortURL{},
	}
}
