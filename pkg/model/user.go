package model

type User struct {
	ID            int64  `db:"id"`
	FirstName     string `db:"first_name"`
	LastName      string `db:"last_name"`
	DateOfBirth   string `db:"date_of_birth"`
	IncomePerYear string `db:"income_per_year"`
}

type JsonUser struct {
	ID            int64  `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	DateOfBirth   string `json:"date_of_birth"`
	IncomePerYear string `json:"income_per_year"`
}

type UserRequest struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	DateOfBirth   string `json:"date_of_birth"`
	IncomePerYear string `json:"income_per_year"`
}
