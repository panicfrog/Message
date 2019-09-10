package data

import (
	"testing"
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
	token := TokenPlayload{
		Account:  "someAccount",
		Platform: PlatformiOS,
	}
	str, err := EncodeToken(&token)
	if err != nil || len(str) == 0 {
		t.Error("encode token 失败")
	}

	token2, err := DecodeToken(str)
	if err != nil || &token2 == nil  {
		t.Error("decode token 失败")
	}

	if !token.Equal(&token2) {
		t.Error("token 验证错误")
	}
}

func TestTokenPlayload_Equal(t *testing.T) {
	t1 := TokenPlayload{
		Account:  "someAccount",
		Platform: PlatformiOS,
	}
	t2 := TokenPlayload{
		Account:  "someAccount",
		Platform: PlatformAndroid,
	}
	t3 := TokenPlayload{
		Account:  "someAccount",
		Platform: PlatformiOS,
	}
	if t1.Equal(&t2) || !t1.Equal(&t3) {
		t.Error("失败")
	}
}