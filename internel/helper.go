package internel

import "time"

// 毫秒
func MicroSec(t time.Time) int64 {
	return t.UnixNano()/1e6
}