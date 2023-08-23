package atomic

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNocmpComparability(t *testing.T) {
	tests := []struct {
		desc       string
		give       interface{}
		comparable bool
	}{
		{
			desc: "nocmp struct",
			give: nocmp{},
		},
		{
			desc: "struct with nocmp embedded",
			give: struct{ nocmp }{},
		},
		{
			desc:       "pointer to struct with nocmp embedded",
			give:       &struct{ nocmp }{},
			comparable: true,
		},

		// All exported types must be uncomparable.
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			typ := reflect.TypeOf(tt.give)
			assert.Equalf(t, tt.comparable, typ.Comparable(),
				"type %v comparablity mismatch", typ)
		})
	}
}
