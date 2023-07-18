package controllers

import (
	"net/http"

	"github.com/Milkado/api-ars-arcanum/database"
	"github.com/Milkado/api-ars-arcanum/models"
	"github.com/gin-gonic/gin"
)

func AllNahelBonds(ctx *gin.Context) {
	var nahelBonds []models.NahelBond
	var nahelBondsTransformed []models.NahelBondsGet
	err := database.DB.Preload("Powers.MagicSystem.Shard").Model(&nahelBonds).Find(&nahelBondsTransformed).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": nahelBondsTransformed})
}

func GetNahelBond(ctx *gin.Context) {
	id := ctx.Param("id")
	var nahelBond models.NahelBond
	var nahelBondsTransformed models.NahelBondsGet
	database.DB.Preload("Powers.MagicSystem.Shard").Model(&nahelBond).First(&nahelBondsTransformed, id)
	ctx.JSON(http.StatusOK, gin.H{"data": nahelBondsTransformed})
}

func CreateNahelBond(ctx *gin.Context) {
	var nahelBond models.NahelBond

	if err := ctx.ShouldBindJSON(&nahelBond); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidadeNahelBond(&nahelBond); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Preload("Powers").Create(&nahelBond)
	ctx.JSON(http.StatusCreated, gin.H{"data": nahelBond})
}

func UpdateNahelBond(ctx *gin.Context) {
	var nahelBond models.NahelBond

	if err := ctx.ShouldBindJSON(&nahelBond); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidadeNahelBond(&nahelBond); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Preload("Powers").Model(&nahelBond).Updates(nahelBond)
	ctx.JSON(http.StatusOK, gin.H{"data": nahelBond})
}

func DeleteNahelBond(ctx *gin.Context) {
	id := ctx.Param("id")
	var nahelBond models.NahelBond
	database.DB.Delete(&nahelBond, id)
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}

func GetNahelBondTrashed(ctx *gin.Context) {
	var nahelBonds []models.NahelBond
	database.DB.Unscoped().Preload("Powers.MagicSystem.Shard").Where("deleted_at IS NOT NULL").Find(&nahelBonds)
	ctx.JSON(http.StatusOK, gin.H{"data": nahelBonds})
}

func RestoreNahelBond(ctx *gin.Context) {
	id := ctx.Param("id")
	var nahelBond models.NahelBond
	database.DB.Unscoped().Model(&nahelBond).Where("id = ?", id).Update("deleted_at", nil)
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
