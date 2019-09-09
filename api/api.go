package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func SetupApi(port int) {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "hello, world",
			})
		})
	_ = r.Run("0.0.0.0:" + strconv.Itoa(port))
}