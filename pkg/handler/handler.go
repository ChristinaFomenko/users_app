package handler

import (
	"fmt"
	"github.com/ChristinaFomenko/users_app/pkg/database"
	"github.com/gorilla/mux"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     database.UserDB
}

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Хэллоууууу")
	}
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
	a.Router.HandleFunc("/user/create", a.CreateUserHandler()).Methods("POST")
	a.Router.HandleFunc("/users", a.GetUsersHandler()).Methods("GET")
	a.Router.HandleFunc("/user", a.GetUserByFieldHandler()).Methods("GET")
}
