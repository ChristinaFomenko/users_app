package model

type User struct {
	ID            int64   `db:"id"`
	FirstName     string  `db:"first_name"`
	LastName      string  `db:"last_name"`
	DateOfBirth   int64   `db:"date_of_birth"`
	IncomePerYear float64 `db:"income_per_year"`
}

type JSONUser struct {
	ID            int64   `json:"id,omitempty"`
	FirstName     string  `json:"first_name,omitempty"`
	LastName      string  `json:"last_name,omitempty"`
	DateOfBirth   int64   `json:"date_of_birth,omitempty"`
	IncomePerYear float64 `json:"income_per_year,omitempty"`
}

type UserRequest struct {
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	DateOfBirth   int64   `json:"date_of_birth"`
	IncomePerYear float64 `json:"income_per_year"`
}
