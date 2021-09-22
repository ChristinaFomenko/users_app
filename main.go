package main

import (
	"github.com/ChristinaFomenko/users_app/app"
	"log"
	"net/http"
	"os"
)

func main() {
	app := app.New()

	http.HandleFunc("/", app.Router.ServeHTTP)

	err := http.ListenAndServe(":8000", nil)
	check(err)
}

func check(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
