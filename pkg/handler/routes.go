package handler

import (
	"fmt"
	"net/http"

	"github.com/ChristinaFomenko/users_app/pkg/database"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type App struct {
	Router   *mux.Router
	db       *sqlx.DB
	repoUser database.UserDB
}

func NewApp(db *sqlx.DB, repo database.UserDB) *App {
	app := &App{db: db, repoUser: repo, Router: mux.NewRouter()}

	app.initHandlers()

	return app
}

func (a *App) initHandlers() {
	a.Router.HandleFunc("/", IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/user/create", CreateUserHandler(a.repoUser)).Methods("POST")
	a.Router.HandleFunc("/get_all_users", GetAllUsersHandler(a.repoUser)).Methods("GET")
	a.Router.HandleFunc("/user", GetUserByFieldHandler(a.repoUser)).Methods("GET")
	a.Router.HandleFunc("/delete_all_users", DeleteAllUsersHandler(a.repoUser)).Methods("DELETE")
}

func IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Хэллоууууу")
	}
}
