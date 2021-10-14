package main

import (
	database "github.com/ChristinaFomenko/users_app/pkg/database"
	"github.com/ChristinaFomenko/users_app/pkg/handler"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {

	app := handler.New()
	app.DB = &database.DB{}
	err := app.DB.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer app.DB.Close()

	http.HandleFunc("/", app.Router.ServeHTTP)

	log.Println("App running..", "Server at http://localhost:8001")
	http.ListenAndServe(":8001", nil)
	log.Fatal(err)
}
