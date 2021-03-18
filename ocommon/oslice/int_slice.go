package oslice

import (
	"strconv"
)

func IntSliceToStringSlice(slice []int) []string {
	result := make([]string, len(slice))
	for index, data := range slice {
		result[index] = strconv.Itoa(data)
	}
	return result
}

func IntSliceContains(slice []int, data int) bool {
	for _, item := range slice {
		if item == data {
			return true
		}
	}
	return false
}

func IntSliceContainsAny(slice []int, args ...int) bool {
	if len(args) == 0 {
		return true
	}
	for _, arg := range args {
		if IntSliceContains(slice, arg) {
			return true
		}
	}
	return false
}

func IntSliceContainsAll(slice []int, args ...int) bool {
	if len(args) == 0 {
		return true
	}
	for _, arg := range args {
		if !IntSliceContains(slice, arg) {
			return false
		}
	}
	return true
}

func IntSliceIsEmpty(slice []int) bool {
	return slice == nil || len(slice) == 0
}
