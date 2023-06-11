package routes

import (
	"log"

	"github.com/Milkado/api-ars-arcanum/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/shard", controllers.AllShards)
	r.POST("/shard", controllers.CreateShard)
	r.PATCH("/shard/:id", controllers.UpdateShard)
	r.DELETE("/shard/:id", controllers.DeleteShard)
	r.GET("/shard/:id", controllers.GetShard)
	r.GET("/shard-trashed", controllers.GetShardTrashed)
	r.PATCH("/shard-restore/:id", controllers.RestoreShard)

	r.GET("/magic-system", controllers.AllMagicSystems)
	r.POST("/magic-system", controllers.CreateMagicSystem)
	r.GET("/magic-system/:id", controllers.GetMagicSystem)
	r.PATCH("/magic-system/:id", controllers.UpdateMagicSystem)
	r.DELETE("/magic-system/:id", controllers.DeleteMagicSystem)
	r.GET("/magic-system-trashed", controllers.GetMagicSystemTrashed)
	r.PATCH("/magic-system-restore/:id", controllers.RestoreMagicSystem)

	r.GET("/power", controllers.AllPowers)
	r.POST("/power", controllers.CreatePower)
	r.GET("/power/:id", controllers.GetPower)
	r.PATCH("/power/:id", controllers.UpdatePower)
	r.DELETE("/power/:id", controllers.DeletePower)
	r.GET("/power-trashed", controllers.GetPowerTrashed)
	r.PATCH("/power-restore/:id", controllers.RestorePower)

	r.GET("/nahel-bond", controllers.AllNahelBonds)
	r.POST("/nahel-bond", controllers.CreateNahelBond)
	r.GET("/nahel-bond/:id", controllers.GetNahelBond)
	// r.PATCH("/nahel-bond/:id", controllers.UpdateNahelBond)
	// r.DELETE("/nahel-bond/:id", controllers.DeleteNahelBond)
	// r.GET("/nahel-bond-trashed", controllers.GetNahelBondTrashed)
	// r.PATCH("/nahel-bond-restore/:id", controllers.RestoreNahelBond)
	log.Fatal(r.Run(":8000"))
}
