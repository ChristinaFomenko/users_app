package handler

import (
	"net/http"

	"github.com/ChristinaFomenko/users_app/pkg/database"
	"github.com/ChristinaFomenko/users_app/pkg/model"
	"github.com/bearatol/lg"
)

func GetUserByFieldHandler(repoUser database.UserDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userField, err := repoUser.GetUser()
		if userField == nil {
			lg.Fatalf("Can't get user, err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
		}
		var resp = make([]model.JsonUser, len(userField))
		for i, user := range userField {
			resp[i] = mapUserJSON(user)
		}

		sendResponse(w, r, resp, http.StatusOK)
		lg.Info("Пользователи получены!")
	}
}
