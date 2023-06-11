package database

import (
	"log"
	"strconv"

	"github.com/Milkado/api-ars-arcanum/helpers"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDb() {

	connection := helpers.GetEnv("DB_USERNAME") + ":" + helpers.GetEnv("DB_PASSWORD") + "@tcp(" + helpers.GetEnv("DB_HOST") + ":" + helpers.GetEnv("DB_PORT") + ")/" + helpers.GetEnv("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func Paginate(ctx *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := ctx.Request.URL.Query()

		page, errPage := strconv.Atoi(q.Get("page"))

		if errPage != nil {
			log.Fatal(errPage)
		}

		if page <= 0 {
			page = 1
		}

		pageSize, errPageSize := strconv.Atoi(q.Get("page_size"))

		if errPageSize != nil {
			log.Fatal(errPageSize)
		}

		if pageSize <= 0 {
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		db = db.Offset(offset).Limit(pageSize)
		return db
	}
}
