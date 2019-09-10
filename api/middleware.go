package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"message/data"
	"message/variable"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			log.Println("no token ")
			sendHTTPError(c, http.StatusUnauthorized, "no token")
			c.Abort()
			return
		}

		// TODO: 检查token是否在内存中如果不在 提示token过期 或 不存在

		t, err := data.DecodeToken(token)
		if err != nil {
			sendHTTPError(c, http.StatusUnauthorized, "token 错误")
		}
		c.Set(variable.TOKEN_KEY, t)
	}
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, token, Language, From")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}