package main

import (
	"github.com/nikhosagala/herobio/db"

	"github.com/nikhosagala/herobio/conf"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db.CreateConnection()

	defer db.CloseConnection()

	conf.InitializeEngine()

	conf.InitializeRoutes()

	conf.Run() // listen and serve on 0.0.0.0:8080
}
