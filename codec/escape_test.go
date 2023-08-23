package codec

import "testing"

func TestEscape(t *testing.T) {
	s := "/"
	t.Logf("%s\n", Escape(s))
}
