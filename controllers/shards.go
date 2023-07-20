package controllers

import (
	"net/http"

	"github.com/Milkado/api-ars-arcanum/database"
	"github.com/Milkado/api-ars-arcanum/models"
	"github.com/gin-gonic/gin"
)

var shardModel models.Shard

// AllShards godoc
// @Summary List all shards
// @Description Route to get all existing shards
// @Tags shards
// @Accept  json
// @Produce  json
// @Success 200 {object} models.ShardGet
// @Failure 400 {object} string
// @Router /shards [get]
func AllShards(ctx *gin.Context) {
	var shardTransformed []models.ShardGet
	database.DB.Preload("MagicSystems").Model(&shardModel).Find(&shardTransformed)
	ctx.JSON(http.StatusOK, gin.H{"data": shardTransformed})
}

// GetShard godoc
// @Summary Get shard by id
// @Description Route to get a shard by id
// @Tags shards
// @Accept  json
// @Produce  json
// @Param id path int true "Shard ID"
// @Success 200 {object} models.ShardGet
// @Failure 400 {object} string
// @Router /shards/{id} [get]
func GetShard(ctx *gin.Context) {
	id := ctx.Param("id")
	var shardTransformed models.ShardGet
	database.DB.Preload("MagicSystems").Model(&shardModel).First(&shardTransformed, id)
	if shardTransformed.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Shard not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": shardTransformed})
}

// CreateShard godoc
// @Summary Create a new shard
// @Description Route to create a new shard
// @Tags shards
// @Accept  json
// @Produce  json
// @Param shard body models.ShardPost true "Shard Model"
// @Success 201 {object} models.ShardPost
// @Failure 400 {object} string
// @Router /shards [post]
func CreateShard(ctx *gin.Context) {
	var shardPost models.ShardPost
	if err := ctx.ShouldBindJSON(&shardPost); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidadeShard(&shardPost); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&shardModel).Create(&shardPost)
	ctx.JSON(http.StatusCreated, gin.H{"data": shardPost})
}

// UpdateShard godoc
// @Summary Update a shard
// @Description Route to update a shard
// @Tags shards
// @Accept  json
// @Produce  json
// @Param id path int true "Shard ID"
// @Param shard body models.ShardPost true "Shard Model"
// @Success 200 {object} models.ShardPost
// @Failure 400 {object} string
// @Router /shards/{id} [patch]
func UpdateShard(ctx *gin.Context) {
	var shardPost models.ShardPost
	id := ctx.Param("id")
	database.DB.Model(&shardModel).First(&shardPost, id)
	if shardPost.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Shard not found"})
		return
	}
	if err := ctx.ShouldBindJSON(&shardPost); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidadeShard(&shardPost); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Preload("MagicSystems").Model(&shardModel).Updates(shardPost)
	ctx.JSON(http.StatusOK, gin.H{"data": shardPost})
}

// DeleteShard godoc
// @Summary Delete a shard
// @Description Route to delete a shard
// @Tags shards
// @Accept  json
// @Produce  json
// @Param id path int true "Shard ID"
// @Success 200 {object} bool
// @Failure 400 {object} string
func DeleteShard(ctx *gin.Context) {
	id := ctx.Param("id")
	database.DB.Delete(&shardModel, id)
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}

// GetShardTrashed godoc
// @Summary Get trashed shards
// @Description Route to get all trashed shards
// @Tags shards
// @Accept  json
// @Produce  json
// @Success 200 {object} models.ShardGet
// @Failure 400 {object} string
// @Router /shards/trashed [get]
func GetShardTrashed(ctx *gin.Context) {
	var shardTransformed []models.ShardGet
	database.DB.Unscoped().Preload("MagicSystems").Where("deleted_at IS NOT NULL").Model(&shardModel).Find(&shardTransformed)
	ctx.JSON(http.StatusOK, gin.H{"data": shardTransformed})
}

// RestoreShard godoc
// @Summary Restore a trashed shard
// @Description Route to restore a trashed shard
// @Tags shards
// @Accept  json
// @Produce  json
// @Param id path int true "Shard ID"
// @Success 200 {object} bool
// @Failure 400 {object} string
func RestoreShard(ctx *gin.Context) {
	id := ctx.Param("id")
	database.DB.Unscoped().Model(&shardModel).Where("id = ?", id).Update("deleted_at", nil)
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
