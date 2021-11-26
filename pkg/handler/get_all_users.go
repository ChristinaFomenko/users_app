package handler

import (
	"net/http"

	"github.com/ChristinaFomenko/users_app/pkg/database"
	"github.com/ChristinaFomenko/users_app/pkg/model"
	"github.com/bearatol/lg"
)

// @Summary Get All Users
// @Security ApiKeyAuth
// @Tags getusers
// @Description get all users
// @Accept  json
// @Produce  json
// @Success 200 {array} database.UserDB "ok"
// @Failure 400,404 {object} sendResponse
// @Failure 500 {object} sendResponse
// @Failure default {object} sendResponse
// @Router /get_all_users [get]

func GetAllUsersHandler(repoUser database.UserDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := repoUser.GetAllUsers()
		if err != nil {
			lg.Fatalf("Can't get users, err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError, "")
			return
		}

		var resp = make([]model.JsonUser, len(users))
		for i, user := range users {
			resp[i] = MakeUserJSON(user)
		}

		sendResponse(w, r, resp, http.StatusOK, "")
		lg.Info("Пользователи получены!")
	}
}
