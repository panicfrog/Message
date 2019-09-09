package main

import (
	"fmt"
	"log"
	"message/api"
	"message/internel"
	"message/websocket"
	"net"
)

func main() {
	go websocket.Setup(internel.Configuration.WebsocketPort,
		func(token string) bool {
			return len(token) > 12
		},
		func(ident string, msg string, conn net.Conn) {
			log.Printf("identifier: %s, msg: %s, conn %p", ident, msg, conn)
			err := websocket.SendMsgToId(ident, fmt.Sprintf("我收到了你的 '%s'", msg))
			if err != nil {
				fmt.Println(err)
			}
		},
	)
	api.SetupApi(internel.Configuration.ApiPort)
}