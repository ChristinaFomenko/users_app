package handler

import (
	"github.com/ChristinaFomenko/users_app/pkg/database"
	"github.com/ChristinaFomenko/users_app/pkg/model"
	"log"
	"net/http"
)

func GetUsersHandler(repoUser database.UserDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := repoUser.GetUsers()
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
