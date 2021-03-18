package omap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapStrStrToString(t *testing.T) {
	tests := []struct {
		Param  map[string]string
		Expect string
	}{
		{map[string]string{"1": "1", "3": "3", "2": "2"}, "1+12+23+3"},
		{map[string]string{"1": "1", "2": "2", "3": "3"}, "1+12+23+3"},
		{map[string]string{}, ""},
	}
	for _, test := range tests {
		assert.Equal(t, MapStrStrToString(test.Param), test.Expect)
	}
}

func BenchmarkMapOfStringToString(b *testing.B) {
	ss := map[string]string{"321": "123", "3211": "123", "32111": "123", "32s1": "123", "s321": "123"}
	for i := 0; i < b.N; i++ {
		MapStrStrToString(ss)
	}
}
