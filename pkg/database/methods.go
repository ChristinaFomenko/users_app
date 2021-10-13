package database

import (
	"github.com/ChristinaFomenko/users_app/pkg/model"
)

func (d *DB) CreateUser(u *model.User) error {
	//dateOfBirth := time.Time{}
	res, err := d.db.Exec(insertUserSchema, u.FirstName, u.LastName, u.DateOfBirth, u.IncomePerYear)
	if err != nil {
		return err
	}
	res.LastInsertId()
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

//func (d *DB) DeleteUsers() ([]*model.User, error) {
//	var users []*model.User
//	err := d.db.Select(&users, "DELETE * FROM users")
//	if err != nil {
//		return users, err
//	}
//
//	return users, nil
//}

//	var resp []*model.User
//	err := d.db.QueryRow("SELECT * FROM user WHERE id = $1").Scan(&resp)
//	if err != nil {
//		return nil, err
//	}
//
//	return resp, nil

//func (d *DB) FindByLastName(lastName string) (*model.User, error) {
//	u := &model.User{}
//	if err := d.db.Exec('SELECT * from users WHERE last_name = $1', lastName).Scan(
//		&u.ID,
//		&u.FirstName,
//		&u.LastName,
//		&u.DateOfBirth,
//		&u.IncomePerYear,
//		); err != nil {
//		return nil, err
//	}
//	return u, nil
//}
