package main

import (
	"os"
	"testing"

	"github.com/shubham88fru/degign-patterns-go/models"
)

var testApp application

func TestMain(m *testing.M) {

	testApp = application{
		Models: *models.New(nil),
	}

	os.Exit(m.Run())
}
