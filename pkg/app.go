package pkg

import (
	"github.com/ChristinaFomenko/users_app/pkg/database"
	"github.com/gorilla/mux"
	"net/http"
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
	a.Router.HandleFunc("/user", a.CreateUserHandler()).Methods("POST")
	a.Router.HandleFunc("/user", a.GetUsersHandler()).Methods("GET")
	a.Router.HandleFunc("/user/{id}", a.GetUserByIDHandler().Methods("GET"))
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}
