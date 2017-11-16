package conf

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/nikhosagala/herobio/controllers"
	"github.com/nikhosagala/herobio/utils"
)

var router *gin.Engine

// InitializeRoutes for App
func InitializeRoutes() {

	router = gin.Default()

	router.Use(utils.CORSMiddleware())

	router.Use(static.Serve("/", static.LocalFile("public", true)))

	router.GET("/api/ping", controllers.Pong)

	router.GET("/api/heroes", controllers.GetHeroes)

	router.GET("/api/hero/:id", controllers.GetHero)

	router.NoRoute(controllers.NotFound)
}

// Router ...
func Router() *gin.Engine {
	return router
}

// Run ...
func Run() {
	router.Run()
}
