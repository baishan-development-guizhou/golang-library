package oslice

func StringSliceContains(slice []string, arg string) bool {
	for _, item := range slice {
		if item == arg {
			return true
		}
	}
	return false
}

func StringSliceContainsAny(slice []string, args ...string) bool {
	if len(args) == 0 {
		return true
	}
	for _, arg := range args {
		if StringSliceContains(slice, arg) {
			return true
		}
	}
	return false
}

func StringSliceContainsAll(slice []string, args ...string) bool {
	if len(args) == 0 {
		return true
	}
	for _, arg := range args {
		if !StringSliceContains(slice, arg) {
			return false
		}
	}
	return true
}

func StringSliceIsEmpty(slice []string) bool {
	return slice == nil || len(slice) == 0
}

func StringSliceCopy(slice []string) []string {
	dst := make([]string, len(slice))
	copy(dst, slice)
	return dst
}

func StringSliceIndex(slice []string, arg string) int {
	for index, item := range slice {
		if item == arg {
			return index
		}
	}
	return -1
}
