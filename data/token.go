package data

import (
	"encoding/json"
	"errors"
	"message/internel"
	"strings"
	"time"
)

type PlatformType int

const (
	PlatformiOS PlatformType     = 0x01
	PlatformAndroid PlatformType = 0x02
	PlatformWeb PlatformType     = 0x03
	PlatfromDesktop PlatformType = 0x04
	PlatformUnknow PlatformType  = 0x05
)

func NewPlatfrom(flat string) PlatformType {
	f := strings.ToLower(flat)
	switch f {
	case "ios":
		return PlatformiOS
	case "android":
		return PlatformAndroid
	case "web":
		return PlatformWeb
	case "desktop":
		return PlatfromDesktop
	default:
		return PlatformUnknow
	}
}

type TokenPlayload struct {
	Account string         `json:"account"`
	Platform PlatformType  `json:"platform"`
	CreateTime int64       `json:"create_time"`
}

func NewTokenPlayload (account string, platform PlatformType) TokenPlayload {
	return TokenPlayload{
		Account:    account,
		Platform:   platform,
		CreateTime: internel.MicroSec(time.Now()),
	}
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
	return t.Account == token.Account && t.Platform == token.Platform && t.CreateTime == token.CreateTime
}

func VerficationToken(token string) bool {
	_, err := DecodeToken(token)
	return err != nil
}