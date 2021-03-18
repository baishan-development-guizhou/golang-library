package oslice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntSliceToStringSlice(t *testing.T) {
	tests := []struct {
		Param  []int
		Expect []string
	}{
		{[]int{1, 3, 5}, []string{"1", "3", "5"}},
		{[]int{1, 3, 5, 7}, []string{"1", "3", "5", "7"}},
		{[]int{}, []string{}},
	}
	for _, test := range tests {
		assert.Equal(t, IntSliceToStringSlice(test.Param), test.Expect)
	}
}

func TestIntSliceContains(t *testing.T) {
	tests := []struct {
		Param1 []int
		Param2 int
		Expect bool
	}{
		{[]int{1, 3, 2, 2, 1}, 3, true},
		{[]int{1, 3, 2, 2, 1}, 1, true},
		{[]int{1, 3, 2, 2, 1}, 2, true},
		{[]int{1, 3, 2, 2, 1}, 4, false},
		{[]int{}, 4, false},
	}
	for _, test := range tests {
		assert.Equal(t, IntSliceContains(test.Param1, test.Param2), test.Expect)
	}
}

func TestIntSliceContainsAny(t *testing.T) {
	tests := []struct {
		Param1 []int
		Param2 []int
		Expect bool
	}{
		{[]int{1, 3, 2, 2, 1}, []int{3, 42}, true},
		{[]int{1, 3, 2, 2, 1}, []int{42, 2}, true},
		{[]int{1, 3, 2, 2, 1}, []int{1, 5}, true},
		{[]int{1, 3, 2, 2, 1}, []int{45, 42}, false},
		{[]int{}, []int{}, true},
		{[]int{}, []int{4}, false},
	}
	for _, test := range tests {
		assert.Equal(t, IntSliceContainsAny(test.Param1, test.Param2...), test.Expect)
	}
}

func TestIntSliceContainsAll(t *testing.T) {
	tests := []struct {
		Param1 []int
		Param2 []int
		Expect bool
	}{
		{[]int{1, 3, 2, 2, 1}, []int{3, 42}, false},
		{[]int{1, 3, 2, 2, 1}, []int{1, 2}, true},
		{[]int{1, 3, 2, 2, 1}, []int{1, 5}, false},
		{[]int{1, 3, 2, 2, 1}, []int{45, 42}, false},
		{[]int{}, []int{}, true},
		{[]int{4}, []int{4}, true},
	}
	for _, test := range tests {
		assert.Equal(t, IntSliceContainsAll(test.Param1, test.Param2...), test.Expect)
	}
}

func TestIntSliceIsEmpty(t *testing.T) {
	tests := []struct {
		Param  []int
		Expect bool
	}{
		{[]int{1, 3, 2, 2, 1}, false},
		{nil, true},
		{[]int{1, 3, 2, 2, 1}, false},
		{[]int{1, 3, 2, 2, 1}, false},
		{[]int{}, true},
		{[]int{4}, false},
	}
	for _, test := range tests {
		assert.Equal(t, IntSliceIsEmpty(test.Param), test.Expect)
	}
}
