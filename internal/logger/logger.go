package logger

import (
	"log"
	"os"
)

type Log struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func NewLog() *Log {
	return &Log{
		InfoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
