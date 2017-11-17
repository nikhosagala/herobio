package conf

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/nikhosagala/herobio/controllers"
)

var router *gin.Engine

// InitializeEngine is function to contruct Gin engine
func InitializeEngine() {

	router = gin.Default()

	config := cors.DefaultConfig()

	config.AllowOrigins = []string{"https://herobio.herokuapp.com/"}

	router.Use(cors.New(config))

	router.Use(gzip.Gzip(gzip.BestCompression))

	router.Use(static.Serve("/", static.LocalFile("public", true)))
}

// InitializeRoutes function to define routing
func InitializeRoutes() {

	router.GET("/api/ping", controllers.Index)

	router.GET("/api/heroes", controllers.GetHeroes)

	router.GET("/api/hero/:id", controllers.GetHero)

	router.NoRoute(controllers.NotFound)
}

// Router return Gin engine
func Router() *gin.Engine {
	return router
}

// Run function to run Gin engine
func Run() {
	router.Run()
}
