package main

import (
	"fmt"
	"log"
	"message/api"
	"message/dbOps"
	"message/internel"
	"message/storage"
	"message/websocket"
)

func main() {
	dbOps.ConfigDB()
	storage.SetupRedis()
	go websocket.Setup(internel.Configuration.WebsocketPort,
		func(token string) bool {
			return len(token) > 12
		},
		func(iden string) {
			log.Println(iden, "连接了")
		},
		func(ident string, msg string) {
			log.Printf("identifier: %s, msg: %s, conn %p", ident, msg)
			err := websocket.SendMsgToId(ident, fmt.Sprintf("我收到了你的 '%s'", msg))
			if err != nil {
				fmt.Println(err)
			}
		},
		func(ident string) {
			log.Println(ident, "关闭了")
		},
	)
	api.SetupApi(internel.Configuration.ApiPort)

}