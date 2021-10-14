package handler

import (
	"encoding/json"
	"github.com/ChristinaFomenko/users_app/pkg/model"
	"log"
	"net/http"
)

func parse(r *http.Request, data interface{}) error {
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

//TODO: это конструктор. "конструкторы в go" погуглить надо. Ты возвращаешь не указатель, значит должно называться MakeUserJSON. Однако, если ты возвращаешь не указатель, то копируешь данные в новый указатель, что не оптимально.
func mapUserJSON(u *model.User) model.JsonUser {
	return model.JsonUser{
		ID:            u.ID,
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		DateOfBirth:   u.DateOfBirth,
		IncomePerYear: u.IncomePerYear,
	}
}
