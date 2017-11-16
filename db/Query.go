package db

import (
	"github.com/nikhosagala/herobio/models"
)

// GetAllHero ...
func GetAllHero(limit int, offset int) []models.Hero {
	heroes := []models.Hero{}
	DB.Find(&heroes).Limit(limit)
	return heroes
}

// GetHero ...
func GetHero(id int) models.Hero {
	hero := models.Hero{}
	DB.Find(&hero, id)
	return hero
}
