package database

import (
	"github.com/ChristinaFomenko/users_app/pkg/model"
)

func (d *DB) CreateUser(u *model.User) error {
	res, err := d.db.Exec(insertUserSchema, u.FirstName, u.LastName, u.DateOfBirth, u.IncomePerYear)
	if err != nil {
		return err
	}
	res.LastInsertId()
	return err
}

func (d *DB) GetUser() ([]*model.User, error) {
	var user []*model.User
	err := d.db.Select(&user, "SELECT * FROM users")
	if err != nil {
		return user, err
	}

	return user, nil
}

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
