package handler

import (
	"github.com/ChristinaFomenko/users_app/pkg/model"
	"log"
	"net/http"
)

func (a *App) GetUserByFieldHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := a.DB.GetUser()
		if userId == nil {
			log.Printf("Can't get user, err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
		}
		var resp = make([]model.JsonUser, len(userId))
		for i, user := range userId {
			resp[i] = mapUserJSON(user)
		}

		sendResponse(w, r, resp, http.StatusOK)
	}
}

func (d *DB) GetUser() ([]*model.User, error) {
	var names []*model.User
	err := d.db.Select(&names, "SELECT first_name, last_name, date_of_birth FROM users LIMIT 2")
	//err := d.db.Select(&names, "SELECT first_name FROM users LIMIT 2")
	if err != nil {
		return nil, err
	}
	return names, err
}
