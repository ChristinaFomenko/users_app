package handler

import (
	"github.com/ChristinaFomenko/users_app/pkg/model"
	"github.com/jmoiron/sqlx"
	"log"
	"math"
	"net/http"
)

type DB struct {
	db *sqlx.DB
}

type User struct {
	ID            int64   `db:"id"`
	FirstName     string  `db:"first_name"`
	LastName      string  `db:"last_name"`
	DateOfBirth   int64   `db:"date_of_birth"`
	IncomePerYear float64 `db:"income_per_year"`
}

type JsonUser struct {
	ID            int64   `json:"id,omitempty"`
	FirstName     string  `json:"first_name,omitempty"`
	LastName      string  `json:"last_name,omitempty"`
	DateOfBirth   int64   `json:"date_of_birth,omitempty"`
	IncomePerYear float64 `json:"income_per_year,omitempty"`
}

func (a *App) CreateUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := model.UserRequest{}
		err := parse(r, &req)
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

//func (d *DB) CreateUser(u *model.User) error {
//	res, err := d.db.Exec(u.FirstName, u.LastName, u.DateOfBirth, u.IncomePerYear)
//	if err != nil {
//		return err
//	}
//	res.LastInsertId() //TODO: функция не void, у нее есть возвращаемые занчения, но они не обрабатываются никак
//	return err
//}
