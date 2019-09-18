package main

import (
	"message/api"
	"message/chat"
	"message/dbOps"
	"message/internel"
	"message/storage"
	"message/websocket"
)

func main() {
	dbOps.ConfigDB()
	storage.SetupRedis()
	go websocket.Setup(internel.Configuration.WebsocketPort,
		storage.VerficationToken,
		chat.AddUserToken,
		chat.DealMessage,
		chat.RemoveUserToken,
	)
	api.SetupApi(internel.Configuration.ApiPort)
}