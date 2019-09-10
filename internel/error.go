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