package body

import "strings"

type Restart struct {
}

func (v Restart) Check() bool {
	return true
}

func (v Restart) ApiCmd() string {
	return strings.ToLower("restart")
}
