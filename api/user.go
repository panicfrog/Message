package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"message/data"
	"message/dbOps"
	"message/internel"
	"message/storage"
	"message/variable"
)

type UserParams struct {
	Account string `json:"account" binding:"required"`
	Passwd string  `json:"passwd" binding:"required"`
}

// 注册
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

// 登录
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

// 获取所有好友
func getFriends(c *gin.Context) {
	token, err := getToken(c)
	if err != nil {
		sendFail(c, err.Error())
		return
	}
	users, err := dbOps.AllFriends(token.Account)
	if err != nil {
		sendFail(c, err.Error())
	}
	sendSuccess(c, "成功", users)
}

// 获取所在所有房间
func getRooms(c *gin.Context) {
	token, err := getToken(c)
	if err != nil {
		sendFail(c, err.Error())
		return
	}
	rooms, err := dbOps.AllRooms(token.Account)
	if err != nil {
		sendFail(c, err.Error())
	}
	sendSuccess(c, "成功", rooms)
}


func getToken(c *gin.Context) (data.TokenPlayload, error) {
	var token data.TokenPlayload
	t, ok := c.Get(variable.TOKEN_KEY)
	if !ok {
		return token, errors.New("")
	}
	token, ok = t.(data.TokenPlayload)
	if !ok {
		return token, errors.New("token 过期或不存在")
	}
	return token, nil
}