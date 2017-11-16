package utils

import (
	"fmt"
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

// CORSMiddleware for Gin
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

// Router ...
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

// LimitAndOffset ...
func LimitAndOffset(limit string, offset string) (int, int) {
	return string2Int(limit), string2Int(offset)
}

// ParseParam2Int ...
func ParseParam2Int(value string) int {
	parse, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return 0
	}
	return int(parse)
}
