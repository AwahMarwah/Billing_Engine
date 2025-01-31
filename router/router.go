package router

import (
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
	return router.Run()
}
