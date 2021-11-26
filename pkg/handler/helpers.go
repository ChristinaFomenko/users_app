package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ChristinaFomenko/users_app/pkg/model"
	"github.com/bearatol/lg"
)

func parse(r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func sendResponse(w http.ResponseWriter, _ *http.Request, data interface{}, status int, s string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		lg.Infof("cant format json, err+%v\n", err)
	}
}

func MakeUserJSON(u *model.User) model.JsonUser {
	return model.JsonUser{
		ID:            u.ID,
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		DateOfBirth:   u.DateOfBirth,
		IncomePerYear: u.IncomePerYear,
	}
}
