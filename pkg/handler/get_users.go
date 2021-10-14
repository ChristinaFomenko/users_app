package handler

import (
	"github.com/ChristinaFomenko/users_app/pkg/model"
	"log"
	"net/http"
)

func (a *App) GetUsersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := GetUsers()
		if err != nil {
			log.Printf("Can't get users, err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]model.JsonUser, len(users))
		for i, user := range users {
			resp[i] = mapUserJSON(user)
		}

		sendResponse(w, r, resp, http.StatusOK)
	}
}

func (d *DB) GetUsers() ([]*model.User, error) {
	var users []*model.User
	err := d.db.Select(&users, "SELECT * FROM users")
	if err != nil {
		return users, err
	}

	return users, nil
}
