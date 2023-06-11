package main

import (
	"fmt"

	"github.com/Milkado/api-ars-arcanum/database"
	"github.com/Milkado/api-ars-arcanum/models"
)

func main() {
	database.ConnectDb()
	fmt.Println("Migrating...")
	database.DB.AutoMigrate(
		&models.Shard{}, &models.MagicSystem{}, &models.Power{}, &models.NahelBond{}, &models.AllomanticTable{},
	)
	fmt.Println("Migration complete!")
}