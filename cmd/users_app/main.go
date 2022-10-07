package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ChristinaFomenko/users_app/pkg/database"
	"github.com/ChristinaFomenko/users_app/pkg/handler"
	"github.com/bearatol/lg"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// @title Users App API
// @version 1.0
// @description API Server for UsersList Application

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	lg.Trace("Connected to database!")
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s "+
		"sslmode=disable", database.DBHost, database.DBPort, database.DBUsername, database.DBTable, database.DBPassword))
	if err != nil {
		lg.Fatal(err)
	}
	defer db.Close()

	db.MustExec(database.CreateSchema)

	repoUser := database.NewUserRepository(db)
	app := handler.NewApp(db, repoUser)

	http.HandleFunc("/", app.Router.ServeHTTP)

	server := &http.Server{
		Addr:              ":8000",
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           app.Router,
	}

	err = server.ListenAndServe()
	if err != nil {
		lg.Fatalf("Listen and serve failed", err)
	}

	lg.Info("App running...", "Server at http://localhost:8000")
}
