package pkg

import (
	"encoding/json"
	"github.com/ChristinaFomenko/users_app/pkg/model"
	"log"
	"net/http"
)

func parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func sendResponse(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("cant format json, err+%v\n", err)
	}
}

func mapUserJSON(u *model.User) model.JsonUser {
	return model.JsonUser{
		ID:            u.ID,
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		DateOfBirth:   u.DateOfBirth,
		IncomePerYear: u.IncomePerYear,
	}
}
