package golang

import (
	"bytes"
	"runtime"
	"strconv"
)

// GetGoroutineID 获取协程id
// golang本身在设计上不存在协程id，在某些情况下存在获取协程id的需求
func GetGoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
