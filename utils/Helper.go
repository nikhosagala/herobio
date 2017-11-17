package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Render content
func Render(c *gin.Context, data gin.H) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.JSON(http.StatusOK, data["payload"])
	}
}

// SetUserStatus for App
func SetUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("token"); err == nil || token != "" {
			c.Set("is_logged_in", true)
		} else {
			c.Set("is_logged_in", false)
		}
	}
}

// Router return new Gin engine
func Router() *gin.Engine {
	return gin.New()
}

func string2Int(value string) int {
	parse, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return 0
	}
	return int(parse)
}

// LimitAndOffset function to parse limit and offset
func LimitAndOffset(limit string, offset string) (int, int) {
	return string2Int(limit), string2Int(offset)
}

// ParseParam2Int parse value to int
func ParseParam2Int(value string) int {
	parse, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return 0
	}
	return int(parse)
}
