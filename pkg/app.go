//TODO: обычно как раз в handlers описываются роуты, а уже в отдельных файлах их обработка.
package pkg

import (
	"github.com/ChristinaFomenko/users_app/pkg/database"
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
	a.Router.HandleFunc("/user/create", a.CreateUserHandler()).Methods("POST")
	a.Router.HandleFunc("/users", a.GetUsersHandler()).Methods("GET")
	a.Router.HandleFunc("/user", a.GetUserByIDHandler()).Methods("GET")
	//a.Router.HandleFunc("/delete/users", a.DeleteUsersHandler()).Methods("DELETE")
}

// Get wraps the router for GET method
//func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
//	a.Router.HandleFunc(path, f).Methods("GET")
//}
