package data

import (
	"message/internel"
	"testing"
	"time"
)

func TestEncodeMessage(t *testing.T) {
	var jsonStr = `{"from": "from","to": "to","type": 1,"content": "内容"}`
	msg, err := EncodeMessage(jsonStr)
	if err != nil {
		t.Error("encodeMessage fail")
	}
	t.Logf("%v", msg)
}

func TestDecodeMessage(t *testing.T) {
	msg := Message{
		From:       "from",
		To:         "to",
		Type:       0,
		Content:    "内容",
	}
	message, err := DecodeMessage(msg)
	if err != nil {
		t.Error("decodeMessage fail")
	}
	t.Log(message)
}

func TestEncodeAndDecodeToken(t *testing.T) {
	token := TokenPlayload {
		Account:  "yeyongping",
		Platform: PlatformiOS,
		CreateTime: internel.MicroSec(time.Now()),
	}
	str, err := EncodeToken(&token)
	if err != nil || len(str) == 0 {
		t.Error("encode token 失败")
	}
	t.Log(str)
	t.Log(token)

	token2, err := DecodeToken(str)
	if err != nil || &token2 == nil  {
		t.Error("decode token 失败")
	}
	t.Log(token2)

	if !token.Equal(&token2) {
		t.Error("token 验证错误")
	}
}

func TestTokenPlayload_Equal(t *testing.T) {
	ti, err :=  time.Parse("2006-01-02 15:04:05", "2019-09-10 15:26:51")
	if err != nil {
		t.Errorf("解析时间错误: %v", err)
	}
	micro := internel.MicroSec(ti)
	t1 := TokenPlayload{
		Account:  "someAccount",
		Platform: PlatformiOS,
		CreateTime: micro,
	}
	t2 := TokenPlayload{
		Account:  "someAccount",
		Platform: PlatformAndroid,
		CreateTime: micro,
	}
	t3 := TokenPlayload{
		Account:  "someAccount",
		Platform: PlatformiOS,
		CreateTime: micro,
	}
	if t1.Equal(&t2) || !t1.Equal(&t3) {
		t.Error("失败")
	}
}