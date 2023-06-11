package seeds

import (
	"fmt"

	"github.com/Milkado/api-ars-arcanum/database"
	"github.com/Milkado/api-ars-arcanum/models"
)

func SeedMagicSystems() {
	fmt.Println("Seeding Magic System...")

	magics := [][]string{
		{"Allomancy", "Harmony"},
		{"Feruchemy", "Harmony"},
		{"Hemalurgy", "Harmony"},
		{"AonDor", "Devotion"},
		{"Dakhor", "Dominion"},
		{"Surgebinding", "Honor"},
		{"Voidbinding", "Odium"},
		{"BioChroma", "Endowment"},
		{"Sand Mastery", "Autonomy"},
		{"ChayShan", "Ambition"},
		{"Awakening", "Endowment"},
	}

	for _, magic := range magics {
		shard := models.Shard{}
		database.DB.Where("name = ?", magic[1]).First(&shard)

		magicSystem := models.MagicSystem{Name: magic[0], ShardID: shard.ID}
		database.DB.Create(&magicSystem)
	}

	fmt.Println("Magic System seeded!")
}
