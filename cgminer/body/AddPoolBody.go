package body

import "strings"

type AddPool struct {
}

func (v AddPool) Check() bool {
	return true
}

func (v AddPool) ApiCmd() string {
	return strings.ToLower("addpool")
}
