package service

import (
	"bookstore-go/model"
	"bookstore-go/repository"
	"encoding/base64"
	"errors"
	"fmt"
)

type UserService struct {
	userDao *repository.UserDao
}

func NewUserService(userDao *repository.UserDao) *UserService {
	return &UserService{
		userDao: userDao,
	}
}

// UserRegister user register
func (u *UserService) UserRegister(username, password, phone, email string) error {
	// step 1. verify the uniqueness of the username, number and email address
	exists, err := u.checkUserExists(username, phone, email)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("username or phone or emasil already exists")
	}

	// step 2. password encryption
	encryptedPassword := u.encryptPassword(password)

	// step 3. save user data
	user := &model.User{
		Username: username,
		Password: encryptedPassword,
		Phone:    phone,
		Email:    email,
	}
	return u.userDao.CreateUser(user)
}

func (u *UserService) checkUserExists(username, phone, email string) (bool, error) {
	// check username
	users, err := u.userDao.GetUsersByIdentities(username, phone, email)
	if err != nil {
		if !errors.Is(err, repository.ErrUserNotFound) {
			return false, err
		}
	}
	return len(users) > 0, nil
}

// 简单的加密处理
// todo 做成一个可拓展模块
func (u *UserService) encryptPassword(password string) string {
	return base64.StdEncoding.EncodeToString([]byte(password))
}
