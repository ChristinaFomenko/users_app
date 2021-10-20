package main

import (
	"fmt"
	database "github.com/ChristinaFomenko/users_app/pkg/database"
	"github.com/ChristinaFomenko/users_app/pkg/handler"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {

	app := handler.New()
	app.DB = &database.DB{}
	err := Open()
	if err != nil {
		log.Fatal(err)
	}

	defer Close()

	http.HandleFunc("/", app.Router.ServeHTTP)

	log.Printf("App running...", "Server at http://localhost:8000")
	err = http.ListenAndServe(":8000", nil)
	log.Fatal(err)
}

func Open() error {
	pg, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", database.DbHost, database.DbPort, database.DbUsername, database.DbTable, database.DbPassword))
	if err != nil {
		return err
	}
	log.Println("Connected to database!")

	pg.MustExec(database.CreateSchema)

	return nil
}

func Close() error {
	return Close()
}
