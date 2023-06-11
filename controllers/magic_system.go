package controllers

import (
	"net/http"

	"github.com/Milkado/api-ars-arcanum/database"
	"github.com/Milkado/api-ars-arcanum/models"
	"github.com/gin-gonic/gin"
	"github.com/Milkado/api-ars-arcanum/transformer/magic_system"
)

func AllMagicSystems(ctx *gin.Context) {
	var magicSystems []transformer.MagicSystem
	database.DB.Preload("Shard").Preload("Powers").Find(&magicSystems)
	ctx.JSON(http.StatusOK, gin.H{"data": magicSystems})
}

func GetMagicSystem(ctx *gin.Context) {
	id := ctx.Param("id")
	var magicSystem transformer.MagicSystem
	database.DB.Preload("Shard").First(&magicSystem, id)
	ctx.JSON(http.StatusOK, gin.H{"data": magicSystem})
}

func CreateMagicSystem(ctx *gin.Context) {
	var magicSystem models.MagicSystem
	if err := ctx.ShouldBindJSON(&magicSystem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Preload("Shard").Create(&magicSystem)
	ctx.JSON(http.StatusCreated, gin.H{"data": magicSystem})
}

func UpdateMagicSystem(ctx *gin.Context) {
	var magicSystem models.MagicSystem
	if err := ctx.ShouldBindJSON(&magicSystem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Preload("Shard").Model(&magicSystem).Updates(magicSystem)
	ctx.JSON(http.StatusOK, gin.H{"data": magicSystem})
}

func DeleteMagicSystem(ctx *gin.Context) {
	id := ctx.Param("id")
	database.DB.Delete(&models.MagicSystem{}, id)
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}

func GetMagicSystemTrashed(ctx *gin.Context) {
	var magicSystems []models.MagicSystem
	database.DB.Unscoped().Preload("Shard").Where("deleted_at IS NOT NULL").Find(&magicSystems)
	ctx.JSON(http.StatusOK, gin.H{"data": magicSystems})
}

func RestoreMagicSystem(ctx *gin.Context) {
	id := ctx.Param("id")
	database.DB.Unscoped().Model(&models.MagicSystem{}).Where("id = ?", id).Update("deleted_at", nil)
}
