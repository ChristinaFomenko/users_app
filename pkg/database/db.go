package database

import (
	"fmt"
	"github.com/ChristinaFomenko/users_app/pkg/model"
	"github.com/jmoiron/sqlx"
	"log"
)

//TODO: уж слишком много берет на себя интерфейс. Будет открываться соединение при получения и создании пользователя. А потом? Создастся, к примеру, таблица "организации", будет еще открываться соединение, или интерфейс UserDB будет еще работать и с организациями?  Вынеси отдельно подключение или вообще просто в main воткни. В main как раз и происходит инициализация всего, а потом в других пакетах происходит обработка.

type UserDB interface {
	Open() error
	Close() error
	CreateUser(u *model.User) error
	GetUsers() ([]*model.User, error)
	GetUser() ([]*model.User, error)
}

type DB struct {
	db *sqlx.DB
}

func (d *DB) CreateUser(u *model.User) error {
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

func (d *DB) GetUsers() ([]*model.User, error) {
	var users []*model.User
	err := d.db.Select(&users, "SELECT * FROM users")
	if err != nil {
		return users, err
	}

	return users, nil
}

func (d *DB) GetUser() ([]*model.User, error) {
	var names []*model.User
	err := d.db.Select(&names, "SELECT first_name, last_name, date_of_birth FROM users LIMIT 2")
	//err := d.db.Select(&names, "SELECT first_name FROM users LIMIT 2")
	if err != nil {
		return nil, err
	}
	return names, err
}

func (d *DB) Open() error {
	pg, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", DbHost, DbPort, DbUsername, DbTable, DbPassword))
	if err != nil {
		return err
	}
	log.Println("Connected to database!")

	pg.MustExec(CreateSchema)
	d.db = pg

	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}
