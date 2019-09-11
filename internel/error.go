package internel

import "errors"

var (
	EncodeMessageError = errors.New("encode message error")
	DecodeMessageError = errors.New("decode message error")
)

// DB
var (
	DBErrorExited = errors.New("data is exits")
)

var (
	RedisTokenNotExited = errors.New("token not exits")
	RedisTokenExpire = errors.New("token expire")
)