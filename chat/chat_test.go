package chat

import "testing"

func TestDecodeAndEncodeSingleToId(t *testing.T) {
	v, err := encodeSingleToId("yeyongping", "xiaohuang")
	if err != nil {
		t.Error(err)
	}
	t.Log(v)
	u1, u2, err := decodeSingleToId(v)
	t.Log(u1, u2)
	if err != nil {
		t.Error(err)
	}
}