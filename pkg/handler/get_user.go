package handler

import (
	"github.com/ChristinaFomenko/users_app/pkg/database"
	"github.com/ChristinaFomenko/users_app/pkg/model"
	"log"
	"net/http"
)

func GetUserByFieldHandler(repoUser database.UserDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userField, err := repoUser.GetUser()
		if userField == nil {
			log.Printf("Can't get user, err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
		}
		var resp = make([]model.JsonUser, len(userField))
		for i, user := range userField {
			resp[i] = mapUserJSON(user)
		}

		sendResponse(w, r, resp, http.StatusOK)
	}
}
