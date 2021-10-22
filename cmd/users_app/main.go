package main

import (
	"fmt"
	database "github.com/ChristinaFomenko/users_app/pkg/database"
	"github.com/ChristinaFomenko/users_app/pkg/handler"
	"github.com/bearatol/lg"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {

	lg.Trace("Connected to database!")
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", database.DbHost, database.DbPort, database.DbUsername, database.DbTable, database.DbPassword))
	if err != nil {
		lg.Fatal(err)
	}
	defer db.Close()

	db.MustExec(database.CreateSchema)

	repoUser := database.NewUserRepository(db)
	app := handler.NewApp(db, repoUser)

	http.HandleFunc("/", app.Router.ServeHTTP)

	lg.Info("App running...", "Server at http://localhost:8000")
	lg.Fatal(http.ListenAndServe(":8000", app.Router))
}
