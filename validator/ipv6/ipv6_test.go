package ipv6

import (
	"testing"
)

func TestValidateString(t *testing.T) {
	ips := []string{
		":::",
		"::1",
		"0:0:0:0:0:0:192.168.12.1",
	}
	for _, ip := range ips {
		isValid, err := ValidateString(ip)
		if isValid {
			t.Logf("Success ip %s valid", ip)
		} else {
			t.Errorf("Fail ip %s err:%s", ip, err)
		}
	}
}
