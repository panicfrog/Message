package data

import (
	"encoding/json"
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
	return internel.AESEncrypt([]byte(internel.Configuration.AesKey), string(b))
}

func DecodeToken(t string) (TokenPlayload, error) {
	var token TokenPlayload
	dt, err := internel.AESDecrypt([]byte(internel.Configuration.AesKey), t)
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