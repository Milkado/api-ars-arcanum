package routes

import (
	"log"

	"github.com/Milkado/api-ars-arcanum/controllers"
	"github.com/Milkado/api-ars-arcanum/middleware"
	"github.com/gin-gonic/gin"
	docs "github.com/Milkado/api-ars-arcanum/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))


	HandleOpen(r)
	HandleAuthed(r)

	log.Fatal(r.Run(":8000"))
}

func HandleAuthed(inc *gin.Engine) {
	inc.Use(middleware.Authenticated())
	inc.GET("/user", controllers.AllUsers)

	inc.POST("/shard", controllers.CreateShard)
	inc.PATCH("/shard/:id", controllers.UpdateShard)
	inc.DELETE("/shard/:id", controllers.DeleteShard)
	inc.GET("/shard-trashed", controllers.GetShardTrashed)
	inc.PATCH("/shard-restore/:id", controllers.RestoreShard)

	inc.POST("/magic-system", controllers.CreateMagicSystem)
	inc.PATCH("/magic-system/:id", controllers.UpdateMagicSystem)
	inc.DELETE("/magic-system/:id", controllers.DeleteMagicSystem)
	inc.GET("/magic-system-trashed", controllers.GetMagicSystemTrashed)
	inc.PATCH("/magic-system-restore/:id", controllers.RestoreMagicSystem)

	inc.POST("/power", controllers.CreatePower)
	inc.PATCH("/power/:id", controllers.UpdatePower)
	inc.DELETE("/power/:id", controllers.DeletePower)
	inc.GET("/power-trashed", controllers.GetPowerTrashed)
	inc.PATCH("/power-restore/:id", controllers.RestorePower)

	inc.POST("/nahel-bond", controllers.CreateNahelBond)
	inc.PATCH("/nahel-bond/:id", controllers.UpdateNahelBond)
	inc.DELETE("/nahel-bond/:id", controllers.DeleteNahelBond)
	inc.GET("/nahel-bond-trashed", controllers.GetNahelBondTrashed)
	inc.PATCH("/nahel-bond-restore/:id", controllers.RestoreNahelBond)

	inc.POST("/allomantic-table", controllers.CreateAllomanticTable)
	inc.PATCH("/allomantic-table/:id", controllers.UpdateAllomanticTable)
	inc.DELETE("/allomantic-table/:id", controllers.DeleteAllomanticTable)
	inc.GET("/allomantic-table-trashed", controllers.GetAllomanticTableTrashed)
	inc.PATCH("/allomantic-table-restore/:id", controllers.RestoreAllomanticTable)
}

func HandleOpen(inc *gin.Engine) {
	inc.POST("/login", controllers.Login)

	inc.GET("/shard", controllers.AllShards)
	inc.GET("/shard/:id", controllers.GetShard)

	inc.GET("/magic-system", controllers.AllMagicSystems)
	inc.GET("/magic-system/:id", controllers.GetMagicSystem)

	inc.GET("/power", controllers.AllPowers)
	inc.GET("/power/:id", controllers.GetPower)

	inc.GET("/nahel-bond", controllers.AllNahelBonds)
	inc.GET("/nahel-bond/:id", controllers.GetNahelBond)

	inc.GET("/allomantic-table", controllers.AllAllomanticTables)
	inc.GET("/allomantic-table/:id", controllers.GetAllomanticTable)
}
