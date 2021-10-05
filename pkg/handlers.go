package pkg

import (
	//"errors"
	"fmt"
	"github.com/ChristinaFomenko/users_app/pkg/model"
	"log"
	"math"
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

		if u.FirstName == "" || u.LastName == "" {
			log.Println("Имя/фамилия - обязательны!")
			log.Printf("Не могу сохранить пользователя в базу данных")
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
		if u.DateOfBirth < 1900 || u.DateOfBirth > 2021 {
			log.Println("Диапазаон дат от 1900 до 2021")
			log.Printf("Не могу сохранить пользователя в базу данных")
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
		if u.IncomePerYear == math.Trunc(u.IncomePerYear) {
			log.Printf("Не могу сохранить пользователя в базу данных, число должно быть с плавающей точкой")
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		} else {
			err = a.DB.CreateUser(u)
			resp := mapUserJSON(u)
			sendResponse(w, r, resp, http.StatusOK)
			log.Println("Пользователь успешно сохранен в бд!")
		}
	}
}

//if len(u.FirstName) == 0 || len(u.FirstName) > 40 {
//	log.Println("Имя должно быть в диапозоне от 1-40 символов!")
//	return
//}
//if len(u.LastName) == 0 || len(u.LastName) > 40 {
//	log.Println("Фамилия должна быть в диапозоне от 1-40 символов!")
//}

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
