package pkg

import (
	//"errors"
	"fmt"
	"github.com/ChristinaFomenko/users_app/pkg/model"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	//"github.com/go-playground/validator/v10"
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
		//datestring := "07-20-2018"
		//fmt.Println(datestring)
		//date, err := time.Parse("01-02-2006", datestring)
		//fmt.Println(date, err)
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
		}
		resp := mapUserJSON(u)
		sendResponse(w, r, resp, http.StatusOK)

		//} else if u.FirstName == "" {
		//	log.Println("name", "The name is required!")
		//}
		//// check the name field is between 1 to 120 chars
		//if len(u.FirstName) == 0 || len(u.FirstName) > 40 {
		//	log.Println("name", "The name field must be between 1-40 chars!")
		//}
		//if u.LastName == "" {
		//	log.Println("last_name", "The last name field is required!")
		//}
		//if len(u.LastName) == 0 || len(u.LastName) > 40 {
		//	log.Println("last name", "The name field must be between 1-40 chars!")
		//}

	}
}

func (a *App) GetUsersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := a.DB.GetUsers()
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

func (a *App) GetUserByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		userId, err := a.DB.GetUser()
		if userId == nil {
			log.Printf("Can't get users, err=%v \n", err)
			sendResponse(w, r, id, http.StatusInternalServerError)
			return
		}

	}
}
