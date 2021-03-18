package ostring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		Param  string
		Expect bool
	}{
		{"123", false},
		{"", true},
		{"124214124", false},
	}
	for _, test := range tests {
		assert.Equal(t, IsEmpty(test.Param), test.Expect)
	}
}

func TestIsNotEmpty(t *testing.T) {
	tests := []struct {
		Param  string
		Expect bool
	}{
		{"123", true},
		{"", false},
		{"124214124", true},
	}
	for _, test := range tests {
		assert.Equal(t, IsNotEmpty(test.Param), test.Expect)
	}
}

func TestIsAllEmpty(t *testing.T) {
	tests := []struct {
		Param  []string
		Expect bool
	}{
		{[]string{"321", "123", "123"}, false},
		{[]string{"321", "", "123"}, false},
		{[]string{"321", "123", ""}, false},
		{[]string{"", "123", "123"}, false},
		{[]string{"", "", ""}, true},
		{[]string{}, true},
	}
	for _, test := range tests {
		assert.Equal(t, IsAllEmpty(test.Param...), test.Expect)
	}
}

func TestIsAnyEmpty(t *testing.T) {
	tests := []struct {
		Param  []string
		Expect bool
	}{
		{[]string{"321", "123", "123"}, false},
		{[]string{"321", "", "123"}, true},
		{[]string{"321", "123", ""}, true},
		{[]string{"", "123", "123"}, true},
		{[]string{"", "", ""}, true},
		{[]string{}, true},
	}
	for _, test := range tests {
		assert.Equal(t, IsAnyEmpty(test.Param...), test.Expect)
	}
}

func TestIsBlank(t *testing.T) {
	tests := []struct {
		Param  string
		Expect bool
	}{
		{"123", false},
		{"      ", true},
		{"124214124", false},
	}
	for _, test := range tests {
		assert.Equal(t, IsBlank(test.Param), test.Expect)
	}
}

func TestIsNotBlank(t *testing.T) {
	tests := []struct {
		Param  string
		Expect bool
	}{
		{"123", true},
		{"      ", false},
		{"124214124", true},
	}
	for _, test := range tests {
		assert.Equal(t, IsNotBlank(test.Param), test.Expect)
	}
}

func TestIsAllBlank(t *testing.T) {
	tests := []struct {
		Param  []string
		Expect bool
	}{
		{[]string{"321", "123", "123"}, false},
		{[]string{"321", " ", "123"}, false},
		{[]string{"321", "123", ""}, false},
		{[]string{"", "123", "123"}, false},
		{[]string{"", " ", ""}, true},
		{[]string{}, true},
	}
	for _, test := range tests {
		assert.Equal(t, IsAllBlank(test.Param...), test.Expect)
	}
}

func TestIsAnyBlank(t *testing.T) {
	tests := []struct {
		Param  []string
		Expect bool
	}{
		{[]string{"321", "123", "123"}, false},
		{[]string{"321", " ", "123"}, true},
		{[]string{"321", "123", " "}, true},
		{[]string{" ", "123", "123"}, true},
		{[]string{" ", " ", ""}, true},
		{[]string{}, true},
	}
	for _, test := range tests {
		assert.Equal(t, IsAnyBlank(test.Param...), test.Expect)
	}
}

func TestDefaultIsEmpty(t *testing.T) {
	tests := []struct {
		Param1 string
		Param2 string
		Expect string
	}{
		{"123", "test", "123"},
		{"", "test", "test"},
		{"124214124", "test", "124214124"},
	}
	for _, test := range tests {
		assert.Equal(t, DefaultIfEmpty(test.Param1, test.Param2), test.Expect)
	}
}
func TestDefaultIsBlank(t *testing.T) {
	tests := []struct {
		Param1 string
		Param2 string
		Expect string
	}{
		{"123", "test", "123"},
		{"   ", "test", "test"},
		{"124214124", "test", "124214124"},
	}
	for _, test := range tests {
		assert.Equal(t, DefaultIfBlank(test.Param1, test.Param2), test.Expect)
	}
}

func TestFirstNotEmpty(t *testing.T) {
	tests := []struct {
		Param1 []string
		Expect string
	}{
		{[]string{"", "", "123"}, "123"},
		{[]string{"   ", "", "123"}, "   "},
		{[]string{"", "", ""}, ""},
		{[]string{}, ""},
	}
	for _, test := range tests {
		assert.Equal(t, FirstNotEmpty(test.Param1...), test.Expect)
	}
}

func TestFirstNotBlank(t *testing.T) {
	tests := []struct {
		Param1 []string
		Expect string
	}{
		{[]string{"", "", "123"}, "123"},
		{[]string{"   ", "", "123"}, "123"},
		{[]string{"", "", ""}, ""},
		{[]string{}, ""},
	}
	for _, test := range tests {
		assert.Equal(t, FirstNotBlank(test.Param1...), test.Expect)
	}
}

func TestFromBytes(t *testing.T) {
	tests := []struct {
		Param1 []byte
		Expect string
	}{
		{[]byte{51, 50, 49}, "321"},
		{[]byte{49, 50, 51}, "123"},
	}
	for _, test := range tests {
		assert.Equal(t, FromBytes(test.Param1), test.Expect)
	}
}

func BenchmarkFromBytes(b *testing.B) {
	var bb = []byte{51, 50, 49}
	for i := 0; i < b.N; i++ {
		FromBytes(bb)
	}
}

func TestToBytes(t *testing.T) {
	ToBytes("321")

	tests := []struct {
		Param1 string
		Expect []byte
	}{
		{"321", []byte{51, 50, 49}},
		{"123", []byte{49, 50, 51}},
	}
	for _, test := range tests {
		assert.Equal(t, ToBytes(test.Param1), test.Expect)
	}
}
func BenchmarkToBytes(b *testing.B) {
	var ss = "321"
	for i := 0; i < b.N; i++ {
		ToBytes(ss)
	}
}
