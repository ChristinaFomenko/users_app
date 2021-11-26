package handler

import (
	"net/http"

	"github.com/ChristinaFomenko/users_app/pkg/database"
	"github.com/ChristinaFomenko/users_app/pkg/model"
	"github.com/bearatol/lg"
)

// @Summary Get User by field
// @Security ApiKeyAuth1
// @Tags getuser
// @Description get user by field
// @Field get-user-by-field
// @Accept  json
// @Produce  json
// @Success 200 {object} database.UserDB
// @Failure 400,404 {object} sendResponse
// @Failure 500 {object} sendResponse
// @Failure default {object} sendResponse
// @Router /user [get]

func GetUserByFieldHandler(repoUser database.UserDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userField, err := repoUser.GetUser()
		if userField == nil {
			lg.Fatalf("Can't get user, err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError, "")
		}
		var resp = make([]model.JsonUser, len(userField))
		for i, user := range userField {
			resp[i] = MakeUserJSON(user)
		}

		sendResponse(w, r, resp, http.StatusOK, "")
		lg.Info("Пользователи получены!")
	}
}
