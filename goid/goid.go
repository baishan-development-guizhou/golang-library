package goid

import "fmt"

func GoID() string {
	return fmt.Sprintf("%v", curGoroutineID())
}
