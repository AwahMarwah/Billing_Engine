package router

import (
	"billing_engine/controller/borrower"
	"billing_engine/controller/health"
	"billing_engine/database"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func Run(db database.DB) (err error) {
	router := gin.Default()
	router.Use(cors.New(corsConfig))
	healthController := health.NewController(db.SqlDb)
	router.GET("health", healthController.Check)
	borrowerGroup := router.Group("borrower")
	{
		borrowerController := borrower.NewController(db.GormDb)
		borrowerGroup.GET("", borrowerController.List)
		borrowerGroup.POST("", borrowerController.Create)
		borrowerGroup.GET("/:id", borrowerController.Detail)
		borrowerGroup.DELETE("/:id", borrowerController.Delete)
	}
	return router.Run()
}
