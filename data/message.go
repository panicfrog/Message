package data

import (
	"encoding/json"
	"log"
	"message/internel"
)

type MessageType int

const (
	TextMessage    MessageType = 0x01
	PictureMessage MessageType = 0x02
	VoiceMessage   MessageType = 0x03
	FileMessage    MessageType = 0x04
)

type Message struct {
	Id          string        `json:"id"`
	From        string        `json:"from"`
	To          string        `json:"to"`
	CreateTime  int64         `json:"create_time"`
	Type        MessageType   `json:"type"`
	Content     string        `json:"content"`
}

func EncodeMessage(msg string) (Message, error) {
	var message Message
	if err := json.Unmarshal([]byte(msg), &message); err != nil {
		log.Println(err)
		return message, internel.EncodeMessageError
	}
	return  message,nil
}

func DecodeMessage(message Message) (string, error) {
	b, err := json.Marshal(&message);
	if err != nil {
		log.Println(err)
		return "", internel.DecodeMessageError
	}
	return string(b), nil
}