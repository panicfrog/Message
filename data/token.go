package data

import (
	"encoding/json"
	"errors"
	"message/internel"
)

type PlatformType int

const (
	PlatformiOS PlatformType     = 0x01
	PlatformAndroid PlatformType = 0x02
	PlatformWeb PlatformType     = 0x03
	PlatfromDesktop PlatformType = 0x04
)

type TokenPlayload struct {
	Account string         `json:"account"`
	Platform PlatformType  `json:"platform"`
}

func EncodeToken(t *TokenPlayload) (string, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return  "", err
	}
	str, err := internel.AESEncrypt([]byte(internel.Configuration.AesKey), string(b))
	if err != nil {
		return "", err
	}
	bt := internel.Base64Encode([]byte(str))
	if len(bt) == 0 {
		return "", errors.New("EncodeToken: Base64Encode error")
	}
	return string(bt), nil
}

func DecodeToken(t string) (TokenPlayload, error) {
	var token TokenPlayload
	bt, err := internel.Base64Decode(t)
	if err != nil {
		return token, err
	}
	dt, err := internel.AESDecrypt([]byte(internel.Configuration.AesKey), string(bt))
	if err != nil {
		return token, err
	}
	err = json.Unmarshal([]byte(dt), &token)
	if err != nil {
		return token, err
	}
	return token, nil
}

func (t *TokenPlayload) Equal(token *TokenPlayload) bool {
	return t.Account == token.Account && t.Platform == token.Platform
}