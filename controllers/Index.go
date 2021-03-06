package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikhosagala/herobio/models"
	"github.com/nikhosagala/herobio/utils"
)

var payload gin.H

// Index to return 
func Index(c *gin.Context) {
	utils.Render(c, gin.H{
		"payload": models.Response{Code: http.StatusOK, Data: "API Running"},
	})
}

// NotFound is function to return if routes not found
func NotFound(c *gin.Context) {
	utils.Render(c, gin.H{
		"payload": models.Response{Code: http.StatusNotFound, Data: []models.Response{}, Error: "Not found"},
	})
}
