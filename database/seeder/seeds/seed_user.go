package seeds

import (
	"fmt"

	"github.com/Milkado/api-ars-arcanum/database"
	"github.com/Milkado/api-ars-arcanum/helpers"
	"github.com/Milkado/api-ars-arcanum/models"
)

func SeedUser() {
	fmt.Println(helpers.Yellow + "Seeding users..." + helpers.Reset)
	hasehdPassword, _ := helpers.HashPassword("1234")
	var users = []models.User{
		{Username: "Khriss", Email: "Khrissalla@gmail.com", Password: hasehdPassword, Name: "Khrissalla"},
	}

	for _, user := range users {
		database.DB.Create(&user)
	}

	fmt.Println(helpers.Green + "Users seeded!" + helpers.Reset)

}