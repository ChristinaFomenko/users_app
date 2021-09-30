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

//func ValidateInputs(data interface{}) (bool, map[string]string) {
//	var validate *validator.Validate
//
//	validate = validator.New()
//
//	err := validate.Struct(data)
//
//	if err != nil {
//		if err, ok := err.(*validator.InvalidValidationError); ok {
//			panic(err)
//		}
//
//		errors := make(map[string]string)
//
//		reflected := reflect.ValueOf(data)
//
//		for _, err := range err.(validator.ValidationErrors) {
//
//			//attempt to find field by name and get json tag name
//			field, _ := reflected.Type().FieldByName(err.StructField())
//			var name string
//
//			if name = field.Tag.Get("json"); name == "" {
//				name = strings.ToLower(err.StructField())
//			}
//
//			switch err.Tag() {
//			case "required":
//				errors[name] = "The " + name + " is required!"
//				break
//			case "field":
//				errors[name] = "The " + name + " should be equal to the " + err.Param()
//				break
//			default:
//				errors[name] = "The " + name + " is invalid"
//				break
//			}
//		}
//
//		return false, errors
//	}
//
//	return true, nil
//}

func mapUserJSON(u *model.User) model.JsonUser {
	return model.JsonUser{
		ID:            u.ID,
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		DateOfBirth:   u.DateOfBirth,
		IncomePerYear: u.IncomePerYear,
	}
}
