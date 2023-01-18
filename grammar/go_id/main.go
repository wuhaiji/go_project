package go_id

import (
	"bytes"
	"runtime"
	"strconv"
)

func ExtractGID(s []byte) int64 {
	s = s[len("goroutine "):]
	s = s[:bytes.IndexByte(s, ' ')]
	gid, _ := strconv.ParseInt(string(s), 10, 64)
	return gid
}

// GetGoroutineID Parse the go id from runtime.Stack() output. Slow, but it works.
func GetGoroutineID() int64 {
	var buf [64]byte
	stack := runtime.Stack(buf[:], false)
	s := buf[:stack]
	return ExtractGID(s)
}
