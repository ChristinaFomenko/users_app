package database

import (
	"github.com/ChristinaFomenko/users_app/pkg/model"
	"github.com/jmoiron/sqlx"
)

type UserDB interface {
	CreateUser(u *model.User) error
	GetUsers() ([]*model.User, error)
	GetUser() ([]*model.User, error)
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

func (d *UserRepository) GetUsers() ([]*model.User, error) {
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
	//err := d.db.Select(&names, "SELECT first_name FROM users LIMIT 2")
	if err != nil {
		return nil, err
	}
	return names, err
}

//func (d *db) Open() error {
//	pg, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", DbHost, DbPort, DbUsername, DbTable, DbPassword))
//	if err != nil {
//		return err
//	}
//	log.Println("Connected to database!")
//
//	pg.MustExec(CreateSchema)
//	d.db = pg
//
//	return nil
//}

//func (d *db) Close() error {
//	return d.db.Close()
//}
