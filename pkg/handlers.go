package pkg

import (
	//"errors"
	"fmt"
	"github.com/ChristinaFomenko/users_app/pkg/model"
	"log"
	"net/http"
	//"github.com/go-playground/validator/v10"
)

type JsonUser struct {
	ID            int64   `json:"id"`
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	DateOfBirth   int64   `json:"date_of_birth"`
	IncomePerYear float64 `json:"income_per_year"`
}

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Хэллоууууу")
	}
}

//func (u *JsonUser) Validate() (map[string] interface{}, bool) {
//	if len(u.FirstName) <= 0 {
//		return nil, false
//	}
//	return nil, false
//
//}

func (a *App) CreateUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := model.UserRequest{}
		err := parse(w, r, &req)
		if err != nil {
			log.Printf("Cannot parse user err=%v \n", err)
			sendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		u := &model.User{
			ID:            0,
			FirstName:     req.FirstName,
			LastName:      req.LastName,
			DateOfBirth:   req.DateOfBirth,
			IncomePerYear: req.IncomePerYear,
		}

		if u.FirstName == "" {
			log.Println("Имя - обязательно!")
			log.Printf("Не могу сохранить пользователя в базу данных")
			sendResponse(w, r, nil, http.StatusInternalServerError)
		}
		//check the name field is between 1 to 120 chars
		if len(u.FirstName) == 0 || len(u.FirstName) > 40 {
			log.Println("Имя должно быть в диапозоне от 1-40 символов!")
			return
		}
		if u.LastName == "" {
			log.Println("Фамилия - обязательна!")
			log.Printf("Не могу сохранить пользователя в базу данных")
			sendResponse(w, r, nil, http.StatusInternalServerError)
		}
		if len(u.LastName) == 0 || len(u.LastName) > 40 {
			log.Println("Фамилия должна быть в диапозоне от 1-40 символов!")
			return
		} else {
			err = a.DB.CreateUser(u)
			resp := mapUserJSON(u)
			sendResponse(w, r, resp, http.StatusOK)
			log.Println("Пользователь успешно сохранен в бд!")
		}
	}
}

func (a *App) GetUsersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := a.DB.GetUsers()
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

//func (a *App) GetUserByIDHandler() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		vars := mux.Vars(r)
//		id := vars["id"]
//
//		userId, err := a.DB.GetUser()
//		if userId == nil {
//			log.Printf("Can't get users, err=%v \n", err)
//			sendResponse(w, r, id, http.StatusInternalServerError)
//		}
//
//	}
//}
