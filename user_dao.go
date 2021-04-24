package main

import (
	"database/sql"
	"github.com/pkg/errors"
	"homework_week2/db"
	"homework_week2/model"
)

type UserDao struct {
	db *sql.DB
}

func NewUserDao() *UserDao {
	return &UserDao{db: db.GetDB()}
}

// GetUserById Dao层方法
func (userDao *UserDao) GetUserById(id int) (*model.User, error) {
	row := userDao.db.QueryRow("select id, nickname, email, password from `user` where id = ?", id)
	if row.Err() != nil {
		return nil, errors.Wrap(row.Err(), "query failed!")
	}
	var user = model.User{}
	err := row.Scan(&user.Id, &user.Nickname, &user.Email, &user.Password)
	if err != nil {
		return nil, errors.Wrap(err, "scan failed!")
	}
	return &user, nil
}
