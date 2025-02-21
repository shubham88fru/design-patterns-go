package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

const port = ":4000"

type application struct {
	templateMap map[string]*template.Template
	config appConfig

}

type appConfig struct {
	useCache bool
}

func main() {
	app := application{
		templateMap: make(map[string]*template.Template),
	}

	flag.BoolVar(&app.config.useCache, "cache", false, "Use a template cache")
	flag.Parse()

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