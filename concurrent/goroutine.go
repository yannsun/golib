package concurrent

import (
	"bytes"
	"runtime"
	"strconv"
)

// GoroutineID return goroutine id, this is only for debug, never use it in production.
func GoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	// if error, just return 0
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
