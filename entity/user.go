package entity

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Fullname  string    `json:"fullname" gorm:"type:varchar(255);not null"`
	Email     string    `json:"email" gorm:"type:varchar(255);not null"`
	Password  string    `json:"-" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegister struct {
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	UserID  int    `json:"user_id"`
	Message string `json:"message"`
}

type UserCookieResponse struct {
	Message string `json:"message"`
}

func NewUserResponse(id int, msg string) UserResponse {
	return UserResponse{
		UserID:  id,
		Message: msg,
	}
}

func NewCookieResponse(msg string) UserCookieResponse {
	return UserCookieResponse{
		Message: msg,
	}
}
