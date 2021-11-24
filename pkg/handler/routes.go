package handler

import (
	"fmt"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

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

	//a.Router.Get("/user", handler)

	a.Router.HandleFunc("/", IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/user/create", CreateUserHandler(a.repoUser)).Methods("POST")
	a.Router.HandleFunc("/get_all_users", GetAllUsersHandler(a.repoUser)).Methods("GET")
	a.Router.HandleFunc("/user", GetUserByFieldHandler(a.repoUser)).Methods("GET")
	a.Router.HandleFunc("/delete_all_users", DeleteAllUsersHandler(a.repoUser)).Methods("DELETE")

	a.Router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:1323/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

}

func IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Хэллоууууу")
	}
}
