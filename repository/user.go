package repository

import (
	"bookstore-go/model"
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

const (
	userTableName = "user"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}

func (u *UserDao) CreateUser(user *model.User) error {
	if user.CreateAt.IsZero() {
		user.CreateAt = time.Now()
	}
	if user.UpdateAt.IsZero() {
		user.UpdateAt = time.Now()
	}
	return u.db.Table(userTableName).Create(user).Error
}

func (u *UserDao) GetUserByName(username string) (*model.User, error) {
	var user model.User
	if err := u.db.Table(userTableName).Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (u *UserDao) GetUserByPhone(phone string) (*model.User, error) {
	var user model.User
	if err := u.db.Table(userTableName).Where("phone = ?", phone).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (u *UserDao) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := u.db.Table(userTableName).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (u *UserDao) GetUsersByIdentities(username, phone, email string) ([]model.User, error) {
	var users []model.User
	err := u.db.Table(userTableName).
		Where("username = ?", username).
		Or("phone = ?", phone).
		Or("email = ?", email).
		Find(&users).Error

	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return users, ErrUserNotFound
	}

	return users, nil
}

func (u *UserDao) ListUsers() ([]*model.User, error) {
	var users []*model.User
	if err := u.db.Table(userTableName).Find(&users).Error; err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, ErrUserNotFound
	}

	return users, nil
}
