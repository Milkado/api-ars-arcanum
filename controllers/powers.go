package controllers

import (
	"net/http"

	"github.com/Milkado/api-ars-arcanum/database"
	"github.com/Milkado/api-ars-arcanum/models"
	"github.com/gin-gonic/gin"
)

func AllPowers(ctx *gin.Context) {
	var powers []models.Power
	var powerTransformed []models.PowerGet
	database.DB.Preload("MagicSystem.Shard").Model(&powers).Find(&powerTransformed)
	ctx.JSON(http.StatusOK, gin.H{"data": powerTransformed})
}

func GetPower(ctx *gin.Context) {
	id := ctx.Param("id")
	var power models.Power
	var powerTransformed models.PowerGet
	database.DB.Preload("MagicSystem.Shard").Model(&power).First(&powerTransformed, id)
	ctx.JSON(http.StatusOK, gin.H{"data": powerTransformed})
}

func CreatePower(ctx *gin.Context) {
	var power models.Power
	if err := ctx.ShouldBindJSON(&power); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidadePower(&power); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Preload("MagicSystem").Create(&power)
	ctx.JSON(http.StatusCreated, gin.H{"data": power})
}

func UpdatePower(ctx *gin.Context) {
	var power models.Power
	if err := ctx.ShouldBindJSON(&power); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidadePower(&power); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Preload("MagicSystem").Model(&power).Updates(power)
	ctx.JSON(http.StatusOK, gin.H{"data": power})
}

func DeletePower(ctx *gin.Context) {
	id := ctx.Param("id")
	database.DB.Delete(&models.Power{}, id)
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}

func GetPowerTrashed(ctx *gin.Context) {
	var powers []models.Power
	database.DB.Unscoped().Preload("MagicSystem").Where("deleted_at IS NOT NULL").Find(&powers)
	ctx.JSON(http.StatusOK, gin.H{"data": powers})
}

func RestorePower(ctx *gin.Context) {
	id := ctx.Param("id")
	database.DB.Unscoped().Model(&models.Power{}).Where("id = ?", id).Update("deleted_at", nil)
}
