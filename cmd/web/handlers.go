package main

import "net/http"

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home!"))
}	