package models

import (
	"github.com/golang-jwt/jwt/v5"
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `json:"username" validate:"nonzero"`
	Password     string `json:"password" validate:"nonzero,min=6,max=16"`
	Email        string `json:"email" gorm:"unique" validate:"nonzero"`
	Name         string `json:"name" validate:"nonzero"`
	Token        string `json:"token" gorm:"-"`
	RefreshToken string `json:"refresh_token" gorm:"-"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func ValidateUser(user *User) error {
	if err := validator.Validate(user); err != nil {
		return err
	}

	return nil
}
