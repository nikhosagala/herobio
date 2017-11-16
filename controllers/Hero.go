package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhosagala/herobio/db"
	"github.com/nikhosagala/herobio/utils"
)

// GetHeroes ...
func GetHeroes(c *gin.Context) {
	limit, offset := utils.LimitAndOffset(c.Query("limit"), c.Query("offset"))
	payload = utils.Success(limit, offset, db.GetAllHero(limit, offset))
	utils.Render(c, payload)
}

// GetHero ...
func GetHero(c *gin.Context) {
	limit, offset := utils.LimitAndOffset(c.Query("limit"), c.Query("offset"))
	id := utils.ParseParam2Int(c.Param("id"))
	hero := db.GetHero(id)
	if hero.ID == 0 {
		payload = utils.NotFountPayload("Hero")
	} else {
		payload = utils.Success(limit, offset, db.GetHero(id))
	}
	utils.Render(c, payload)
}
