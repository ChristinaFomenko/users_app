package model

type User struct {
	ID            int64   `db:"id"`
	FirstName     string  `db:"first_name" validate:"required,min=1"`
	LastName      string  `db:"last_name" validate:"required,min=1"`
	DateOfBirth   int64   `db:"date_of_birth" validate:"required"`
	IncomePerYear float64 `db:"income_per_year" validate:"required"`
}

type JsonUser struct {
	ID            int64   `json:"id"`
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	DateOfBirth   int64   `json:"date_of_birth"`
	IncomePerYear float64 `json:"income_per_year"`
}

type UserRequest struct {
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	DateOfBirth   int64   `json:"date_of_birth"`
	IncomePerYear float64 `json:"income_per_year"`
}
