package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/shubham88fru/degign-patterns-go/models"
)

const port = ":4000"

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
	DB          *sql.DB
	Models      models.Models
}

type appConfig struct {
	useCache bool
	dsn      string
}

func main() {
	app := application{
		templateMap: make(map[string]*template.Template),
	}

	flag.BoolVar(&app.config.useCache, "cache", false, "Use a template cache")
	flag.StringVar(&app.config.dsn, "dsn", "mariadb:myverysecretpassword@tcp(localhost:3306)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s", "DSN")
	flag.Parse()

	//connect to db
	db, err := initMySqlDB(app.config.dsn)
	if err != nil {
		log.Panic(err)
	}
	app.DB = db
	app.Models = *models.New(db)

	srv := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	fmt.Println("Starting server on port", port)
	err = srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}
