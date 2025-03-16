package main

import (
	"os"
	"testing"

	"github.com/shubham88fru/degign-patterns-go/adapters"
	"github.com/shubham88fru/degign-patterns-go/configuration"
)

var testApp application

func TestMain(m *testing.M) {
	testBackend := &adapters.TestBackend{}
	testAdapter := &adapters.RemoteService{Remote: testBackend}

	testApp = application{
		App: configuration.New(nil, testAdapter),
	}

	os.Exit(m.Run())
}
