package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikhosagala/herobio/models"
)

// NotFountPayload return json payload for gin.H
func NotFountPayload(what string) gin.H {
	return gin.H{"payload": models.Response{Code: http.StatusNotFound, Data: []models.Response{}, Error: what + " not found"}}
}

// Success return defaul json payload for response 
func Success(limit int, offset int, data interface{}) gin.H {
	return gin.H{"payload": models.Response{Code: http.StatusOK, Data: data, Limit: limit, Offset: offset, Error: ""}}
}
