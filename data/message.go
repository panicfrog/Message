package data

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
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
	gorm.Model
	From        string        `gorm:"type:varchar(50);not null" json:"from"`
	To          string        `gorm:"type:varchar(100);not null" json:"to"`
	Type        MessageType   `gorm:"type:integer" json:"type"`
	Content     string        `gorm:"type:varchar(500);not null" json:"content"`
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