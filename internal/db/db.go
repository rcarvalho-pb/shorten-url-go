package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rcarvalho-pb/link-shorter/internal/logger"
)

var myLog *logger.Log

func InitDB() *sql.DB {
	myLog = logger.NewLog()

	db := connDB()
	if db == nil {
		myLog.ErrorLog.Panic("Can't connect to DB")
	}

	return db
}

func connDB() *sql.DB {

	count := 0
	for {
		db, err := openDB()
		if err != nil {
			myLog.InfoLog.Println("DB not ready...")
		} else {
			myLog.InfoLog.Println("Connected to the DB")
			return db
		}

		if count > 10 {
			return nil
		}

		myLog.InfoLog.Println("Backing off for 1 sec...")
		count++
		time.Sleep(1 * time.Second)
	}
}

func openDB() (*sql.DB, error) {
	dsn := os.Getenv("DSN")
	log.Println(dsn)
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
