package body

import "strings"

type Check struct {
	Exists string `json:"Exists"`
	Access string `json:"Access"`
}

func (v Check) Check() bool {
	return true
}

func (v Check) ApiCmd() string {
	return strings.ToLower("Check")
}
