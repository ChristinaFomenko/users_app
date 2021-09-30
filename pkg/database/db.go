package database

import (
	"github.com/ChristinaFomenko/users_app/pkg/model"
	"github.com/jmoiron/sqlx"
	"log"
)

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

func (d *DB) Open() error {
	pg, err := sqlx.Open("postgres", pgConnStr)
	if err != nil {
		return err
	}
	log.Println("Connected to database!")

	pg.MustExec(createSchema)
	d.db = pg

	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}
