package chat

import (
	"errors"
	"fmt"
	"log"
	"message/data"
	"message/dbOps"
	"message/internel"
	"message/websocket"
	"sync"
)

var userMapper sync.Map

func init() {
	userMapper = sync.Map{}
}

func sendToUser(account, message string) error {
	users, ok := userMapper.Load(account)
	if ok {
		us, _ok := users.([]string)
		if !_ok {
			return errors.New("类型错误")
		}
		for _, u := range us {
			if err := websocket.SendMsgToId(u, message); err != nil {
				log.Println(err)
			}
		}
	}
	return nil
}

func SendToUser(account, message string) error {
	// TODO decode message

	// 在线消息
	if err := sendToUser(account, message); err != nil {
		return err
	}
	// TODO 缓存消息

	return nil
}

func SendToRoom(room, message string) error {
	// TODO decode message

	// TODO verify room
	r, err := dbOps.RoomInfo(room)
	if err != nil {
		return err
	}

	for _, u := range r.Users {
		if err := sendToUser(u.Account, message); err != nil {
			return err
		}
	}

	// TODO 缓存消息
	return nil
}

func DealMessage(identity, message string) {
	log.Printf("identifier: %s, msg: %s", identity, message)
	err := websocket.SendMsgToId(identity, fmt.Sprintf("我收到了你的 '%s'", message))
	if err != nil {
		fmt.Println("发生错误：", err)
	}
}

func AddUserToken(token string)  {
	log.Println(token, "添加了")
	t, err := data.DecodeToken(token)
	if err != nil {
		log.Println(err)
		return
	}
	act, loaded := userMapper.LoadOrStore(t.Account, []string{token})
	if loaded {
		actual, ok := act.([]string)
		if !ok {
			log.Println(errors.New("类型错误"))
			return
		}
		actual = append(actual, token)
		userMapper.Store(t.Account, actual)
	}
}

func RemoveUserToken(token string) {
	log.Println(token, "删除了")
	t, err := data.DecodeToken(token)
	if err != nil {
		log.Println(err)
		return
	}
	act, ok := userMapper.Load(t.Account)
	if !ok {
		log.Println(errors.New("不存在"))
		return
	}
	actual, ok := act.([]string)
	if !ok {
		log.Println(errors.New("类型错误"))
		return
	}
	actual2 := []string{}
	for _, v := range actual {
		if v != token {
			actual2 = append(actual2, v)
		}
	}
	if len(actual) != 0 && len(actual) != len(actual2) + 1 {
		log.Println(internel.ChatTokenNotExited)
		return
	}
}