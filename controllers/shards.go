package controllers

import (
	"net/http"

	"github.com/Milkado/api-ars-arcanum/database"
	"github.com/Milkado/api-ars-arcanum/models"
	"github.com/gin-gonic/gin"
	"github.com/Milkado/api-ars-arcanum/transformer/shard"
)

func AllShards(ctx *gin.Context) {
	var shards []transformer.Shard
	database.DB.Preload("MagicSystems").Find(&shards)
	ctx.JSON(http.StatusOK, gin.H{"data": shards})
}

func GetShard(ctx *gin.Context) {
	id := ctx.Param("id")
	var shard transformer.Shard
	database.DB.Preload("MagicSystems").First(&shard, id)
	ctx.JSON(http.StatusOK, gin.H{"data": shard})
}

func CreateShard(ctx *gin.Context) {
	var shard models.Shard
	if err := ctx.ShouldBindJSON(&shard); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&shard)
	ctx.JSON(http.StatusCreated, gin.H{"data": shard})
}

func UpdateShard(ctx *gin.Context) {
	var shard models.Shard
	if err := ctx.ShouldBindJSON(&shard); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Preload("MagicSystems").Model(&shard).Updates(shard)
	ctx.JSON(http.StatusOK, gin.H{"data": shard})
}

func DeleteShard(ctx *gin.Context) {
	id := ctx.Param("id")
	database.DB.Delete(&models.Shard{}, id)
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}

func GetShardTrashed(ctx *gin.Context) {
	var shards []models.Shard
	database.DB.Unscoped().Preload("MagicSystems").Where("deleted_at IS NOT NULL").Find(&shards)
	ctx.JSON(http.StatusOK, gin.H{"data": shards})
}

func RestoreShard(ctx *gin.Context) {
	id := ctx.Param("id")
	database.DB.Unscoped().Model(&models.Shard{}).Where("id = ?", id).Update("deleted_at", nil)
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
