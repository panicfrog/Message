package api

import (
	"github.com/gin-gonic/gin"
	"message/data"
	"message/dbOps"
	"message/internel"
	"message/storage"
)

type UserParams struct {
	Account string `json:"account" binding:"required"`
	Passwd string  `json:"passwd" binding:"required"`
}

func register(c *gin.Context)  {
	var userParams UserParams
	err := c.ShouldBindJSON(&userParams)
	if err != nil {
		sendParamsError(c, err.Error())
		return
	}
	err = dbOps.AddUser(userParams.Account, userParams.Passwd)
	if err == internel.DBErrorExited {
		sendFail(c, "用户已存在")
	} else {
		sendSuccess(c, "注册成功", nil)
	}
}


func login(c *gin.Context) {
	var userParams UserParams
	err := c.ShouldBindJSON(&userParams)
	if err != nil {
		sendParamsError(c, err.Error())
		return
	}

	user, err := dbOps.VerificationUser(userParams.Account, userParams.Passwd)
	if err != nil && err == internel.DBErrorExited {
		sendFail(c, "密码错误或者用户不存在")
		return
	} else if err != nil {
		sendServerInternelError(c, err.Error())
		return
	}
	p := c.Request.Header.Get("platform")
	platform := data.NewPlatfrom(p)
	if p == "" {
	}
	nt := data.NewTokenPlayload(user.Account, platform)
	token, err := data.EncodeToken(&nt)
	if err != nil {
		sendServerInternelError(c, err.Error())
		return
	}
	err = storage.SetToken(nt)
	if err != nil {
		sendServerInternelError(c, err.Error())
		return
	}
	sendSuccess(c, "登录成功", token)
}