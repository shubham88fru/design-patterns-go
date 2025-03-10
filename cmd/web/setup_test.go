package main

import (
	"log"
	"os"
	"testing"

	"github.com/shubham88fru/degign-patterns-go/models"
)

var testApp application

func TestMain(m *testing.M) {
	dsn := "mariadb:myverysecretpassword@tcp(localhost:3306)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s"
	db, err := initMySqlDB(dsn)

	if err != nil {
		log.Panic(err)
	}

	testApp = application{
		DB:     db,
		Models: *models.New(db),
	}

	os.Exit(m.Run())
}
