package db

import (
	"github.com/nikhosagala/herobio/models"
)

// GetAllHero query to get all Hero from database
// param limit and offset
func GetAllHero(limit int, offset int) []models.Hero {
	heroes := []models.Hero{}
	DB.Find(&heroes).Limit(limit)
	return heroes
}

// GetHero query to get Hero from database by id
func GetHero(id int) models.Hero {
	hero := models.Hero{}
	DB.Find(&hero, id)
	return hero
}
