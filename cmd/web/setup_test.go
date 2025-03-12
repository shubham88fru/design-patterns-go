package main

import (
	"os"
	"testing"

	"github.com/shubham88fru/degign-patterns-go/configuration"
)

var testApp application

func TestMain(m *testing.M) {

	testApp = application{
		App: configuration.New(nil),
	}

	os.Exit(m.Run())
}
