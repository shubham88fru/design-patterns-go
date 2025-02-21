package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const port = ":4000"

type application struct {

}

func main() {
	app := application{}

	srv := &http.Server{
		Addr: port,
		Handler: app.routes(),
		IdleTimeout: 30*time.Second,
		ReadTimeout: 5*time.Second,
		ReadHeaderTimeout: 5*time.Second,
		WriteTimeout: 10*time.Second,
	}
	
	fmt.Println("Starting server on port", port)
	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}