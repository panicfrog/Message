package chat

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"message/data"
	"message/dbOps"
	"message/internel"
	"message/websocket"
	"strings"
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
	var msg data.Message
	err = json.Unmarshal([]byte(message), &msg)
	if err != nil {
		log.Println(err)
		return
	}
	if msg.Mode == data.SingleChat { // 单聊
		// TODO 1.解出聊天的两个人 2.判断是否是好友 3.发送
	} else if msg.Mode == data.GroupChat { // 群聊
		// TODO 1.找出房间 2.验证是否是房间成员 3.发送
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

// TODO add test

func decodeSingleToId(id string) (string, string, error) {
	ss := strings.Split(id, "_")
	if len(ss) != 2 {
		return "", "", errors.New("解析出错")
	}

	s1, err := internel.Base64Decode(ss[0])
	if err != nil {
		return "", "", err
	}
	s2, err := internel.Base64Decode(ss[1])
	if err != nil {
		return "", "", err
	}
	return string(s1), string(s2), nil
}

func encodeSingle(user1, user2 string) (string, error) {
	ss := []string{}
	if user1 > user2 {
		ss = append(ss, user1)
		ss = append(ss, user2)
	} else {
		ss = append(ss, user2)
		ss = append(ss, user1)
	}
	result := []string{}
	for i, s := range ss {
		b := internel.Base64Encode([]byte(s))
		result = append(result,  string(b))
		if i == 0 {
			result = append(result,"_")
		}
	}
	return result[0] + result[1] + result[2], nil
}