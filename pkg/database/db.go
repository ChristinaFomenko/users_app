package database

import (
	"fmt"

	"github.com/ChristinaFomenko/users_app/pkg/model"
	"github.com/jmoiron/sqlx"
)

type UserDB interface {
	CreateUser(u *model.User) error
	GetAllUsers() ([]*model.User, error)
	GetUser() ([]*model.User, error)
	DeleteAllUsers() ([]*model.User, error)
}

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (d *UserRepository) CreateUser(u *model.User) error {
	res, err := d.db.Exec(InsertUserSchema, u.FirstName, u.LastName, u.DateOfBirth, u.IncomePerYear)
	if err != nil {
		return err
	}
	_, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return err
}

func (d *UserRepository) GetAllUsers() ([]*model.User, error) {
	var users []*model.User
	err := d.db.Select(&users, "SELECT * FROM users")
	if err != nil {
		return users, err
	}

	return users, nil
}

func (d *UserRepository) GetUser() ([]*model.User, error) {
	var names []*model.User
	err := d.db.Select(&names, "SELECT first_name, last_name, date_of_birth FROM users LIMIT 2")
	if err != nil {
		return nil, err
	}
	return names, err
}

func (d *UserRepository) DeleteAllUsers() ([]*model.User, error) {
	var users []*model.User
	res, err := d.db.Exec("DELETE FROM users")
	if err != nil {
		return nil, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, fmt.Errorf("db is empty")
	}

	return users, nil
}
