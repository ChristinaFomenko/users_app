package main

import (
	"github.com/ChristinaFomenko/users_app/app"
	"github.com/ChristinaFomenko/users_app/app/database"
	"log"
	"net/http"
	"os"
)

func main() {
	app := app.New()
	app.DB = &database.DB{}
	err := app.DB.Open()
	check(err)

	defer app.DB.Close()

	http.HandleFunc("/", app.Router.ServeHTTP)

	log.Println("App running..")
	err = http.ListenAndServe(":8000", nil)
	check(err)
}

func check(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
