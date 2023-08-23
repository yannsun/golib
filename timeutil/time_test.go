package timeutil

import (
	"testing"
	"time"
)

func TestParseTime(t *testing.T)  {
	var tt time.Time
	tt = time.Now()
	tstr := tt.UTC().Format("2006-01-02T15:04:05Z")
	t.Logf("time %s\n", tstr)
}
