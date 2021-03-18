package ostring

import (
	"reflect"
	"strings"
	"unsafe"
)

const (
	_emptyString = ""
)

//IsEmpty determines whether string is empty
func IsEmpty(arg string) bool {
	return arg == _emptyString
}

//IsNotEmpty determines whether string is not empty
func IsNotEmpty(arg string) bool {
	return arg != _emptyString
}

//IsAllEmpty determines whether all string variable parameters are all empty
func IsAllEmpty(args ...string) bool {
	if len(args) <= 0 {
		return true
	}
	for _, arg := range args {
		if IsNotEmpty(arg) {
			return false
		}
	}
	return true
}

//IsAnyEmpty determines whether string variable parameters in which there is empty
func IsAnyEmpty(args ...string) bool {
	if len(args) <= 0 {
		return true
	}
	for _, arg := range args {
		if IsEmpty(arg) {
			return true
		}
	}
	return false
}

//IsBlank determines whether string is blank
func IsBlank(arg string) bool {
	return IsEmpty(arg) || IsEmpty(strings.TrimSpace(arg))
}

//IsNotBlank determines whether string is not blank
func IsNotBlank(arg string) bool {
	return IsNotEmpty(arg) && IsNotEmpty(strings.TrimSpace(arg))
}

//IsAllBlank determines whether all string variable parameters are all blank
func IsAllBlank(args ...string) bool {
	if len(args) <= 0 {
		return true
	}
	for _, arg := range args {
		if IsNotBlank(arg) {
			return false
		}
	}
	return true
}

//IsAnyBlank determines whether string variable parameters in which there is blank
func IsAnyBlank(args ...string) bool {
	if len(args) <= 0 {
		return true
	}
	for _, arg := range args {
		if IsBlank(arg) {
			return true
		}
	}
	return false
}

//DefaultIfEmpty returns default value if parameter is empty,else parameter
func DefaultIfEmpty(arg string, def string) string {
	if IsEmpty(arg) {
		return def
	}
	return arg
}

//DefaultIfBlank returns default value if parameter is blank,else parameter
func DefaultIfBlank(arg string, def string) string {
	if IsBlank(arg) {
		return def
	}
	return arg
}

//FirstNotEmpty returns the first not empty parameter
func FirstNotEmpty(args ...string) string {
	if len(args) <= 0 {
		return _emptyString
	}
	for _, arg := range args {
		if IsNotEmpty(arg) {
			return arg
		}
	}
	return _emptyString
}

//FirstNotBlank returns the first not blank parameter
func FirstNotBlank(args ...string) string {
	if len(args) <= 0 {
		return _emptyString
	}
	for _, arg := range args {
		if IsNotBlank(arg) {
			return arg
		}
	}
	return _emptyString
}

//FromBytes converts byte slice to string,can only be used in read-only.
func FromBytes(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

//ToBytes converts string to byte slice,can only be used in read-only.
func ToBytes(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := *(*reflect.SliceHeader)(unsafe.Pointer(&s))
	bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
	return b
}
