package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Milkado/api-ars-arcanum/database"
	"github.com/Milkado/api-ars-arcanum/models"
	"github.com/Milkado/api-ars-arcanum/transformer/nahel_bond"
	"github.com/gin-gonic/gin"
)

func AllNahelBonds(ctx *gin.Context) {
	var nahelBonds []transformer.NahelBond
	database.DB.Preload("Powers.MagicSystem").Find(&nahelBonds)
	ctx.JSON(http.StatusOK, gin.H{"data": nahelBonds})
}

func GetNahelBond(ctx *gin.Context) {
	id := ctx.Param("id")
	var nahelBond models.NahelBond
	database.DB.Preload("Powers").First(&nahelBond, id)
	ctx.JSON(http.StatusOK, gin.H{"data": nahelBond})
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
