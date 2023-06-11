package seeds

import (
	"fmt"
	"github.com/Milkado/api-ars-arcanum/models"
	"github.com/Milkado/api-ars-arcanum/database"
)


func SeedShards() {
	fmt.Println("Seeding shards...")
	var shards = []models.Shard{
		{ Name: "Harmony", Vessel: "Sazed", Planet: "Scadrial"},
		{ Name: "Devotion", Vessel: "Aona", Planet: "Sel"},
		{ Name: "Dominion", Vessel: "Skai", Planet: "Sel"},
		{ Name: "Preservation", Vessel: "Leras", Planet: "Scadrial"},
		{ Name: "Ruin", Vessel: "Ati", Planet: "Scadrial"},
		{ Name: "Cultivation", Vessel: "Unknown", Planet: "Roshar"},
		{ Name: "Odium", Vessel: "Rayse", Planet: "Braize"},
		{ Name: "Endowment", Vessel: "Edgli", Planet: "Nalthis"},
		{ Name: "Autonomy", Vessel: "Bavadin", Planet: "Taldain"},
		{ Name: "Ambition", Vessel: "Uli Da", Planet: "Threnody"},
		{ Name: "Honor", Vessel: "Tanavast", Planet: "Roshar"},
		{ Name: "Whimsy", Vessel: "Aona", Planet: "Unknown"},
		{ Name: "Mercy", Vessel: "Skai", Planet: "Unknown"},
		{ Name: "Valor", Vessel: "Unknown", Planet: "Unknown"},
		{ Name: "Invention", Vessel: "Unknown", Planet: "Unknown"},
	}

	for _, shard := range shards {
		database.DB.Create(&shard)
	}
	fmt.Println("Shards seeded!")
}