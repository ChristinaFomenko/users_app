package handler

import (
	"math"
	"net/http"

	"github.com/ChristinaFomenko/users_app/pkg/database"
	"github.com/ChristinaFomenko/users_app/pkg/model"
	"github.com/bearatol/lg"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	db *sqlx.DB
}

type User struct {
	ID            int64   `repoUser:"id"`
	FirstName     string  `repoUser:"first_name"`
	LastName      string  `repoUser:"last_name"`
	DateOfBirth   int64   `repoUser:"date_of_birth"`
	IncomePerYear float64 `repoUser:"income_per_year"`
}

type JsonUser struct {
	ID            int64   `json:"id,omitempty"`
	FirstName     string  `json:"first_name,omitempty"`
	LastName      string  `json:"last_name,omitempty"`
	DateOfBirth   int64   `json:"date_of_birth,omitempty"`
	IncomePerYear float64 `json:"income_per_year,omitempty"`
}

// @Summary Create User
// @Security ApiKeyAuth
// @Tags create
// @Description create user
// @Accept  json
// @Produce  json
// @Param input body user.User "user info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} sendResponse
// @Failure 500 {object} sendResponse
// @Failure default {object} sendResponse
// @Router /user/create [post]

func CreateUserHandler(repoUser database.UserDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := model.UserRequest{}
		err := parse(r, &req)
		if err != nil {
			lg.Infof("Cannot parse user err=%v \n", err)
			sendResponse(w, r, nil, http.StatusBadRequest, "invalid input body")
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
			lg.Fatal("Не могу сохранить пользователя в базу данных. Имя/фамилия - обязательны!")
			sendResponse(w, r, nil, http.StatusInternalServerError, "invalid input body")
			return
		}
		if u.DateOfBirth < 1900 || u.DateOfBirth > 2021 {
			lg.Fatal("Не могу сохранить пользователя в базу данных. Диапазон дат от 1900 до 2021")
			sendResponse(w, r, nil, http.StatusInternalServerError, "invalid input body")
			return
		}
		if u.IncomePerYear == math.Trunc(u.IncomePerYear) {
			lg.Fatal("Не могу сохранить пользователя в базу данных, число должно быть с плавающей точкой")
			sendResponse(w, r, nil, http.StatusInternalServerError, "invalid input body")
			return
		} else {
			err = repoUser.CreateUser(u)
			resp := MakeUserJSON(u)
			sendResponse(w, r, resp, http.StatusOK, "")
			lg.Info("Пользователь успешно сохранен в бд!")
		}
	}
}
