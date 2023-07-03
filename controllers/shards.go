package controllers

import (
	"net/http"

	"github.com/Milkado/api-ars-arcanum/database"
	"github.com/Milkado/api-ars-arcanum/models"
	"github.com/gin-gonic/gin"
)

func AllShards(ctx *gin.Context) {
	var shards []models.Shard
	var shardTransformed []models.ShardReturn
	database.DB.Preload("MagicSystems").Model(&shards).Find(&shardTransformed)
	ctx.JSON(http.StatusOK, gin.H{"data": shardTransformed})
}

func GetShard(ctx *gin.Context) {
	id := ctx.Param("id")
	var shard models.Shard
	var shardTransformed models.ShardReturn
	database.DB.Preload("MagicSystems").Model(&shard).First(&shardTransformed, id)
	if shard.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Shard not found"})
		return
	}
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
	id := ctx.Param("id")
	database.DB.First(&shard, id)
	if shard.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Shard not found"})
		return
	}
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
