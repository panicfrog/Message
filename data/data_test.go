package data

import (
	"message/internel"
	"testing"
	"time"
)

func TestEncodeMessage(t *testing.T) {
	var jsonStr = `{"id": "djadjsajoifwefdakl","from": "from","to": "to","create_time": 1568008420379,"type": 1,"content": "内容"}`
	msg, err := EncodeMessage(jsonStr)
	if err != nil {
		t.Error("encodeMessage fail")
	}
	t.Logf("%v", msg)
}

func TestDecodeMessage(t *testing.T) {
	msg := Message{
		Id:         "djadjsajoifwefdakl",
		From:       "from",
		To:         "to",
		CreateTime: internel.MicroSec(time.Now()),
		Type:       0,
		Content:    "内容",
	}
	message, err := DecodeMessage(msg)
	if err != nil {
		t.Error("decodeMessage fail")
	}
	t.Log(message)
}