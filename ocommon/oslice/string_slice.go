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
