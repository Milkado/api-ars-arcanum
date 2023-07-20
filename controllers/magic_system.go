package controllers

import (
	"net/http"

	"github.com/Milkado/api-ars-arcanum/database"
	"github.com/Milkado/api-ars-arcanum/models"
	"github.com/gin-gonic/gin"
)

var magicSystemModel models.MagicSystem

// AllMagicSystems godoc
// @Summary List all magic systems
// @Description Route to get all existing magic systems
// @Tags magic-systems
// @Accept  json
// @Produce  json
// @Success 200 {object} models.MagicSystemGet
// @Failure 400 {object} string
// @Router /magic-systems [get]
func AllMagicSystems(ctx *gin.Context) {
	var magicSystemTransformed []models.MagicSystemGet
	err := database.DB.Preload("Shard").Model(&magicSystemModel).Find(&magicSystemTransformed).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": magicSystemTransformed})
}

// GetMagicSystem godoc
// @Summary Get magic system by id
// @Description Route to get a magic system by id
// @Tags magic-systems
// @Accept  json
// @Produce  json
// @Param id path int true "Magic System ID"
// @Success 200 {object} models.MagicSystemGet
// @Failure 400 {object} string
// @Router /magic-systems/{id} [get]
func GetMagicSystem(ctx *gin.Context) {
	id := ctx.Param("id")
	var magicSystemTransformed models.MagicSystemGet
	database.DB.Preload("Shard").Model(&magicSystemModel).First(&magicSystemTransformed, id)
	ctx.JSON(http.StatusOK, gin.H{"data": magicSystemTransformed})
}

// CreateMagicSystem godoc
// @Summary Create a new magic system
// @Description Route to create a new magic system
// @Tags magic-systems
// @Accept  json
// @Produce  json
// @Param magicSystem body models.MagicSystemPost true "Magic System Model"
// @Success 201 {object} models.MagicSystemPost
// @Failure 400 {object} string
// @Router /magic-systems [post]
func CreateMagicSystem(ctx *gin.Context) {
	var magicSystemPost models.MagicSystemPost
	if err := ctx.ShouldBindJSON(&magicSystemPost); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidadeMagicSystem(&magicSystemPost); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Preload("Shard").Model(&magicSystemModel).Create(&magicSystemPost)
	ctx.JSON(http.StatusCreated, gin.H{"data": magicSystemPost})
}

// UpdateMagicSystem godoc
// @Summary Update a magic system
// @Description Route to update a magic system
// @Tags magic-systems
// @Accept  json
// @Produce  json
// @Param id path int true "Magic System ID"
// @Param magicSystem body models.MagicSystemPost true "Magic System Model"
// @Success 200 {object} models.MagicSystemPost
// @Failure 400 {object} string
// @Router /magic-systems/{id} [patch]
func UpdateMagicSystem(ctx *gin.Context) {
	var magicSystemPost models.MagicSystemPost
	if err := ctx.ShouldBindJSON(&magicSystemPost); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidadeMagicSystem(&magicSystemPost); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Preload("Shard").Model(&magicSystemModel).Updates(magicSystemPost)
	ctx.JSON(http.StatusOK, gin.H{"data": magicSystemPost})
}

// DeleteMagicSystem godoc
// @Summary Delete a magic system
// @Description Route to delete a magic system
// @Tags magic-systems
// @Accept  json
// @Produce  json
// @Param id path int true "Magic System ID"
// @Success 200 {object} bool
// @Failure 400 {object} string
// @Router /magic-systems/{id} [delete]
func DeleteMagicSystem(ctx *gin.Context) {
	id := ctx.Param("id")
	database.DB.Delete(&magicSystemModel, id)
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}

// GetMagicSystemTrashed godoc
// @Summary List all trashed magic systems
// @Description Route to get all trashed magic systems
// @Tags magic-systems
// @Accept  json
// @Produce  json
// @Success 200 {object} models.MagicSystemGet
// @Failure 400 {object} string
// @Router /magic-systems-trashed [get]
func GetMagicSystemTrashed(ctx *gin.Context) {
	var magicSystemTransformed []models.MagicSystemGet
	database.DB.Unscoped().Preload("Shard").Where("deleted_at IS NOT NULL").Model(&magicSystemModel).Find(&magicSystemTransformed)
	ctx.JSON(http.StatusOK, gin.H{"data": magicSystemTransformed})
}

// RestoreMagicSystem godoc
// @Summary Restore a trashed magic system
// @Description Route to restore a trashed magic system
// @Tags magic-systems
// @Accept  json
// @Produce  json
// @Param id path int true "Magic System ID"
// @Success 200 {object} bool
// @Failure 400 {object} string
// @Router /magic-systems-restore/{id} [patch]
func RestoreMagicSystem(ctx *gin.Context) {
	id := ctx.Param("id")
	database.DB.Unscoped().Model(&magicSystemModel).Where("id = ?", id).Update("deleted_at", nil)
}
