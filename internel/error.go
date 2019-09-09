package internel

import "errors"

var (
	EncodeMessageError = errors.New("encode message error")
	DecodeMessageError = errors.New("decode message error")
)