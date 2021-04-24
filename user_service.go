package main

import (
	"github.com/pkg/errors"
	"homework_week2/model"
)

type UserService struct {
	userDao *UserDao
}

func NewUserService() *UserService {
	return &UserService{userDao: NewUserDao()}
}

// GetUserInfo 根据id查询用户信息
func (us *UserService)GetUserInfo(id int) (*model.User, error) {
	user, err := us.userDao.GetUserById(id)
	if err != nil {
		return nil, errors.WithMessage(err, "get user info failed")
	}
	// 做一些处理。。。
	user.Password = ""
	return user, nil
}