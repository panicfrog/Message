package api

import (
	"github.com/gin-gonic/gin"
	"message/data"
	"message/variable"
	"strconv"
)

func SetupApi(port int) {
	r := gin.Default()

	// common
	setupCommon(r)

	// auth
	authGroup := r.Group("/auth")
	authGroup.Use(AuthMiddleWare())
	setupAuth(authGroup)

	_ = r.Run("0.0.0.0:" + strconv.Itoa(port))
}

func setupCommon(r *gin.Engine) {
	r.GET("/hello", func(c *gin.Context) {
		sendSuccess(c, "success", "world")
	})
}

func setupAuth(g *gin.RouterGroup) {
	g.GET("/token", func(c *gin.Context) {
		t, ok := c.Get(variable.TOKEN_KEY)
		if !ok {
			sendFail(c, "auth失败")
			return
		}
		token, ok := t.(data.TokenPlayload)
		if !ok {
			sendFail(c, "token 出错")
			return
		}
		sendSuccess(c, "成功", token)
	})
}