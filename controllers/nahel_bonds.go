package controllers

import (
	"encoding/json"
	"io"
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

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = json.Unmarshal(body, &nahelBond)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Preload("Powers").Create(&nahelBond)
	ctx.JSON(http.StatusCreated, gin.H{"data": nahelBond})
}
