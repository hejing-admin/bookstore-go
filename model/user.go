package model

import "time"

type User struct {
	ID       int       `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Avatar   string    `json:"avatar"` // 头像
	IsAdmin  bool      `json:"is_admin"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
