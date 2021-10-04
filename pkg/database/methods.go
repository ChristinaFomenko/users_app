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

//func (d *DB) GetUser() ([]*model.User, error) {
//	var user []*model.User
//	err := d.db.Select(&user, 'SELECT * from user WHERE id = $1', )
//}

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
