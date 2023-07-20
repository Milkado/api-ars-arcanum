package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/Milkado/api-ars-arcanum/controllers"
	"github.com/Milkado/api-ars-arcanum/database"
	"github.com/Milkado/api-ars-arcanum/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID uint

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.TestMode)
	routes := gin.Default()

	return routes
}

func CreateShardMock() {
	shard := models.Shard{
		Name:   "Test Shard",
		Vessel: "Test Vessel",
		Planet: "Test Planet",
	}
	err := database.DB.Create(&shard).Error
	if err != nil {
		fmt.Println(err)
	}
	ID = shard.ID
}

func DeleteShardMock() {
	var models models.Shard

	database.DB.Where("id = ?", ID).Delete(&models)
	database.DB.Unscoped().Where("id = ?", ID).Delete(&models)
}

func TestStatusCodeShardGetAll(t *testing.T) {
	database.ConnectDb()
	CreateShardMock()
	defer DeleteShardMock()
	r := SetupTestRoutes()
	r.GET("/shard", controllers.AllShards)
	req, err := http.NewRequest("GET", "/shard", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "Should return 200")
}

//Need fix, mock not working
func TestShardShow(t *testing.T) {
	database.ConnectDb()
	CreateShardMock()
	defer DeleteShardMock()
	r := SetupTestRoutes()
	r.GET("/shard/:id", controllers.GetShard)
	path := "/shard/" + strconv.Itoa(int(ID))
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "Should return 200")
}

func TestShardDelete(t *testing.T) {
	database.ConnectDb()
	CreateShardMock()
	defer database.DB.Unscoped().Where("id = ?", ID).Delete(&models.Shard{})
	r := SetupTestRoutes()
	r.DELETE("/shard/:id", controllers.DeleteShard)
	path := "/shard/" + strconv.Itoa(int(ID))
	req, err := http.NewRequest("DELETE", path, nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "Should return 200")
}

//Need fix, mock not working
func TestShardEdit(t *testing.T) {
	database.ConnectDb()
	CreateShardMock()
	defer DeleteShardMock()
	r := SetupTestRoutes()
	r.PATCH("/shard/:id", controllers.UpdateShard)
	path := "/shard/" + strconv.Itoa(int(ID))
	shard := models.Shard{ 
		Name:   "Test Shard2",
		Vessel: "Test Vessel2",
		Planet: "Test Planet2",
	}
	jsonValue, _ := json.Marshal(shard)
	req, err := http.NewRequest("PATCH", path, bytes.NewBuffer(jsonValue))

	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	var shardMock models.Shard
	json.Unmarshal(resp.Body.Bytes(), &shardMock)
	assert.Equal(t, "Test Vessel2", shardMock.Vessel, "Should return Test Vessel2")
	fmt.Println(shardMock)
}
