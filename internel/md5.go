package internel

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}