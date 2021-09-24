package pkg

import (
	"fmt"
	"github.com/ChristinaFomenko/users_app/pkg/model"
	"log"
	"net/http"
	"time"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Хэллоууууу")
	}
}

func (a *App) CreateUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := model.UserRequest{}
		err := parse(w, r, &req)
		datestring := "07-20-2018"
		fmt.Println(datestring)
		date, err := time.Parse("01-02-2006", datestring)
		fmt.Println(date, err)
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

		err = a.DB.CreateUser(u)
		if err != nil {
			log.Printf("Can't save user to db, err%v\n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := mapUserJSON(u)
		sendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *App) GetUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := a.DB.GetUser()
		if err != nil {
			log.Printf("Can't get users, err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]model.JsonUser, len(user))
		for i, users := range user {
			resp[i] = mapUserJSON(users)
		}

		sendResponse(w, r, resp, http.StatusOK)
	}
}
