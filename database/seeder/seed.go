package main

import (
	"github.com/Milkado/api-ars-arcanum/database"
	"github.com/Milkado/api-ars-arcanum/database/seeder/seeds"
	"gorm.io/gorm"
)

func main() {
	database.ConnectDb()
	database.DB.Transaction(func(tx *gorm.DB) error {
		// seeds.SeedShards()
		// seeds.SeedMagicSystems()
		seeds.SeedUser()
		return nil
	})

}
