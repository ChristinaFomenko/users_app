package handler

import (
	"net/http"

	"github.com/ChristinaFomenko/users_app/pkg/database"
	"github.com/ChristinaFomenko/users_app/pkg/model"
	"github.com/bearatol/lg"
)

// @Summary GetUsersHandler
// @tag.getusers
// @Accept  json
// @Produce  json
// @Param input body "Get Users"
// @Success 200 {object} model.User
// @Failure 400,404 {object} sendResponse
// @Failure 500 {object} sendResponse
// @Failure default {object} sendResponse
// @Router /get_all_users [get]

func GetAllUsersHandler(repoUser database.UserDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := repoUser.GetAllUsers()
		if err != nil {
			lg.Fatalf("Can't get users, err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]model.JsonUser, len(users))
		for i, user := range users {
			resp[i] = mapUserJSON(user)
		}

		sendResponse(w, r, resp, http.StatusOK)
		lg.Info("Пользователи получены!")
	}
}
