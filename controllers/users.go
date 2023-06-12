package controllers

import (
	"net/http"

	"github.com/Milkado/api-ars-arcanum/database"
	"github.com/Milkado/api-ars-arcanum/helpers"
	"github.com/Milkado/api-ars-arcanum/models"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var user models.User
	var loginInput models.LoginInput

	if err := ctx.ShouldBindJSON(&loginInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Where("email = ?", loginInput.Email).First(&user)

	if user.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if helpers.CheckPasswordHash(loginInput.Password, user.Password) == false {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		return
	}

	token, _ := helpers.GenerateToken(user.Email, user.Name, user.ID)

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

// AllUsers returns all users
func AllUsers(ctx *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	ctx.JSON(200, gin.H{"data": users})
}
