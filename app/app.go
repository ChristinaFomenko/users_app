package app

import (
	"github.com/ChristinaFomenko/users_app/app/database"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     database.UserDB
}

func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}

	a.initRoutes()
	return a
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
}
