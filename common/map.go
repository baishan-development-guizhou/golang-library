package common

import (
	"fmt"
	"strings"
)

// MapToString is simple map to properties string.
// eg:
//    { name: "1", age: "1" }  =>  "name=1,age=1"
//
// Tip: The type is float64, when type of int convert to interface, if you want to cast .
func MapToString(mapValues map[string]interface{}) string {
	values := make([]string, 0)
	for key, value := range mapValues {
		values = append(values, fmt.Sprintf("%v=%v", key, value))
	}
	return strings.Join(values, ",")
}
