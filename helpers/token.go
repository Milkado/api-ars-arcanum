package helpers

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type SignedUser struct {
	ID    uint
	Email string
	Name  string
	jwt.RegisteredClaims
}

func GenerateToken(email, name string, id uint) (tokenSigned string, err error) {
	secret := []byte(GetEnv("JWT_SECRET"))

	claims := &SignedUser{
		ID:    id,
		Email: email,
		Name:  name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(GetExpirationTime(10)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, errToken := token.SignedString(secret)

	if errToken != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenSigned string) (err error) {
	token, err := jwt.ParseWithClaims(
		tokenSigned,
		&SignedUser{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(GetEnv("JWT_SECRET")), nil
		},
	)

	if err != nil {
		return err
	}

	claims, ok := token.Claims.(*SignedUser)

	if !ok {
		err = errors.New("Invalid token")
		return
	}

	if claims.ExpiresAt.Unix() < time.Now().Unix() {
		err = errors.New("Expired token")
		return
	}

	return
}
