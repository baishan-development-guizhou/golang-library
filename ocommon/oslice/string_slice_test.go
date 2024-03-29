package oslice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringSliceContains(t *testing.T) {
	tests := []struct {
		Param1 []string
		Param2 string
		Expect bool
	}{
		{[]string{"1", "3", "2", "2", "1"}, "3", true},
		{[]string{"1", "3", "2", "2", "1"}, "1", true},
		{[]string{"1", "3", "2", "2", "1"}, "2", true},
		{[]string{"1", "3", "2", "2", "1"}, "4", false},
		{[]string{}, "4", false},
	}
	for _, test := range tests {
		assert.Equal(t, StringSliceContains(test.Param1, test.Param2), test.Expect)
	}
}

func TestStringSliceContainsAny(t *testing.T) {
	tests := []struct {
		Param1 []string
		Param2 []string
		Expect bool
	}{
		{[]string{"1", "3", "2", "2", "1"}, []string{"3", "42"}, true},
		{[]string{"1", "3", "2", "2", "1"}, []string{"42", "2"}, true},
		{[]string{"1", "3", "2", "2", "1"}, []string{"1", "5"}, true},
		{[]string{"1", "3", "2", "2", "1"}, []string{"45", "42"}, false},
		{[]string{}, []string{}, true},
		{[]string{}, []string{"4"}, false},
	}
	for _, test := range tests {
		assert.Equal(t, StringSliceContainsAny(test.Param1, test.Param2...), test.Expect)
	}
}

func TestStringSliceContainsAll(t *testing.T) {
	tests := []struct {
		Param1 []string
		Param2 []string
		Expect bool
	}{
		{[]string{"1", "3", "2", "2", "1"}, []string{"3", "42"}, false},
		{[]string{"1", "3", "2", "2", "1"}, []string{"1", "2"}, true},
		{[]string{"1", "3", "2", "2", "1"}, []string{"1", "5"}, false},
		{[]string{"1", "3", "2", "2", "1"}, []string{"45", "42"}, false},
		{[]string{}, []string{}, true},
		{[]string{"4"}, []string{"4"}, true},
	}
	for _, test := range tests {
		assert.Equal(t, StringSliceContainsAll(test.Param1, test.Param2...), test.Expect)
	}
}

func TestStringSliceIsEmpty(t *testing.T) {
	tests := []struct {
		Param  []string
		Expect bool
	}{
		{[]string{"1", "3", "2", "2", "1"}, false},
		{nil, true},
		{[]string{"1", "3", "2", "2", "1"}, false},
		{[]string{"1", "3", "2", "2", "1"}, false},
		{[]string{}, true},
		{[]string{"4"}, false},
	}
	for _, test := range tests {
		assert.Equal(t, StringSliceIsEmpty(test.Param), test.Expect)
	}
}

func TestStringSliceCopy(t *testing.T) {
	tests := []struct {
		Param []string
	}{
		{[]string{"1", "3", "2"}},
		{[]string{"1", "2", "4"}},
		{[]string{"1", "3"}},
		{[]string{"1"}},
	}
	for _, test := range tests {
		assert.NotEqual(t, fmt.Sprintf("%p", StringSliceCopy(test.Param)), fmt.Sprintf("%p", test.Param))
	}
}

func TestStringSliceIndex(t *testing.T) {
	tests := []struct {
		Param1 []string
		Param2 string
		Expect int
	}{
		{[]string{"1", "3", "2"}, "2", 2},
		{[]string{"1", "2", "4"}, "1", 0},
		{[]string{"1", "3"}, "3", 1},
		{[]string{"1"}, "2", -1},
	}
	for _, test := range tests {
		assert.Equal(t, StringSliceIndex(test.Param1, test.Param2), test.Expect)
	}
}
