package omap

import (
	"bytes"
	"sort"
)

//MapStrStrToString Converts map string string to string
func MapStrStrToString(data map[string]string) string {

	var result bytes.Buffer
	keys := make([]string, len(data))
	var i = 0
	for key := range data {
		keys[i] = key
		i += 1

	}
	sort.Strings(keys)
	for _, key := range keys {
		result.WriteString(key)
		result.WriteString("+")
		result.WriteString(data[key])
	}
	return result.String()
}
