package concurrent

import "testing"

func TestGoroutineID(t *testing.T) {
	t.Logf("goroutine id :%d", GoroutineID())
}
