package db

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	"github.com/nikhosagala/herobio/models"
)

// DB ...
var DB *gorm.DB

// CreateConnection ...
func CreateConnection() *gorm.DB {
	var url = "host=localhost user=postgres dbname=herobio sslmode=disable password=postgres"
	if strings.Compare(os.Getenv("GIN_MODE"), gin.ReleaseMode) == 0 {
		url = os.Getenv("DATABASE_URL")
	}
	db, err := gorm.Open("postgres", url)

	if err != nil {
		panic("failed to connect database")
	}

	DB = db

	Migrate()

	return db
}

// Migrate ...
func Migrate() {
	DB.DropTableIfExists(&models.Hero{})
	DB.AutoMigrate(&models.Hero{})
}

// CloseConnection ...
func CloseConnection() {
	DB.Close()
}
