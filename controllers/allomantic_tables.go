package controllers

import (
	"net/http"

	"github.com/Milkado/api-ars-arcanum/database"
	"github.com/Milkado/api-ars-arcanum/models"
	"github.com/gin-gonic/gin"
	"github.com/Milkado/api-ars-arcanum/transformer/allomantic_table"
)

func AllAllomanticTables(ctx *gin.Context) {
	var allomanticTables []transformer.AllomanticTable
	database.DB.Preload("Power.MagicSystem").Find(&allomanticTables)
	ctx.JSON(http.StatusOK, gin.H{"data": allomanticTables})
}

func GetAllomanticTable(ctx *gin.Context) {
	id := ctx.Param("id")
	var allomanticTable transformer.AllomanticTable
	database.DB.Preload("Power.MagicSystem").First(&allomanticTable, id)
	ctx.JSON(http.StatusOK, gin.H{"data": allomanticTable})
}

func CreateAllomanticTable(ctx *gin.Context) {
	var allomanticTable models.AllomanticTable
	if err := ctx.ShouldBindJSON(&allomanticTable); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&allomanticTable)
	ctx.JSON(http.StatusCreated, gin.H{"data": allomanticTable})
}

func UpdateAllomanticTable(ctx *gin.Context) {
	var allomanticTable models.AllomanticTable
	if err := ctx.ShouldBindJSON(&allomanticTable); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Preload("Shard").Model(&allomanticTable).Updates(allomanticTable)
	ctx.JSON(http.StatusOK, gin.H{"data": allomanticTable})
}

func DeleteAllomanticTable(ctx *gin.Context) {
	id := ctx.Param("id")
	database.DB.Delete(&models.AllomanticTable{}, id)
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}

func GetAllomanticTableTrashed(ctx *gin.Context) {
	var allomanticTables []models.AllomanticTable
	database.DB.Unscoped().Preload("Power.MagicSystem").Where("deleted_at IS NOT NULL").Find(&allomanticTables)
	ctx.JSON(http.StatusOK, gin.H{"data": allomanticTables})
}

func RestoreAllomanticTable(ctx *gin.Context) {
	id := ctx.Param("id")
	database.DB.Unscoped().Model(&models.AllomanticTable{}).Where("id = ?", id).Update("deleted_at", nil)
}
