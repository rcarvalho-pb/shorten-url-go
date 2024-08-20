package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rcarvalho-pb/link-shorter/internal/logger"
)

var myLog *logger.Log

func Serve() {
	myLog = logger.NewLog()

	port := os.Getenv("PORT")

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: routes(),
	}

	myLog.InfoLog.Println("Starting server...")
	if err := srv.ListenAndServe(); err != nil {
		myLog.ErrorLog.Panic(err)
	}

}
