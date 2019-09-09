package internel

import "testing"


func TestAESDecryptAndEncrypt(t *testing.T) {
	const key = "yeyongping123456"
	message := "测试数据"
	encoded, err := AESEncrypt([]byte(key), message)
	if err != nil {
		t.Error(err)
	}

	msg, err := AESDecrypt([]byte(key), encoded)
	if err != nil {
		t.Error(err)
	}
	if msg != message {
		t.Errorf("AES 加解密错误 '%s' != '%s'", msg, message)
	}
}