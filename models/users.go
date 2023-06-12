package models

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"unique"`
	Name     string `json:"name"`
	Token   string `json:"token" gorm:"-"`
	RefreshToken  string `json:"refresh_token" gorm:"-"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
