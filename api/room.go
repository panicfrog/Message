package api

import (
	"github.com/gin-gonic/gin"
	"message/dbOps"
)

type RoomsParams struct {
	RoomId string	 `json:"roomId" binding:"required"`
}
func getUsersInRoom(c *gin.Context) {
	var userInRoomsParams RoomsParams
	if err := c.ShouldBindJSON(&userInRoomsParams); err != nil {
		sendParamsError(c, err.Error())
		return
	}
	_, err := getToken(c)
	if err != nil {
		sendFail(c, err.Error())
		return
	}
	users, err := dbOps.AllUserInRoom(userInRoomsParams.RoomId)
	if err != nil {
		sendFail(c, err.Error())
		return
	}
	sendSuccess(c, "成功", users)
}

func getInfoOfRoom(c *gin.Context) {
	var userInRoomsParams RoomsParams
	if err := c.ShouldBindJSON(&userInRoomsParams); err != nil {
		sendParamsError(c, err.Error())
		return
	}
	_, err := getToken(c)
	if err != nil {
		sendFail(c, err.Error())
		return
	}
	room, err := dbOps.RoomInfo(userInRoomsParams.RoomId)
	if err != nil {
		sendFail(c, err.Error())
		return
	}
	sendSuccess(c, "成功", room)
}