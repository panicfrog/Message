package data

import (
	"fmt"
	"testing"
	"time"
)

func TestEncodeMessage(t *testing.T) {
	var jsonStr = `{"id": "djadjsajoifwefdakl","from": "from","to": "to","create_time": 1568008131000814000,"type": 1,"content": "内容"}`
	msg, err := EncodeMessage(jsonStr)
	if err != nil {
		t.Error("encodeMessage fail")
	}
	fmt.Printf("%v", msg)
}

func TestDecodeMessage(t *testing.T) {
	msg := Message{
		Id:         "djadjsajoifwefdakl",
		From:       "from",
		To:         "to",
		CreateTime: time.Now().UnixNano(),
		Type:       0,
		Content:    "",
	}
	message, err := DecodeMessage(msg)
	if err != nil {
		t.Error("decodeMessage fail")
	}
	fmt.Println(message)
}