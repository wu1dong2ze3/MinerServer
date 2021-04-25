package body

import (
	"fmt"
	"strings"
)

//{"CGMiner":"2.8.7","API":"1.20"}
type Version struct {
	CGMiner string
	API     string
}

func (v Version) Check() bool {
	fmt.Println("Version:Check:CGMiner=", v.CGMiner, " API=", v.API)
	return true
}
func (v Version) ApiCmd() string {
	return strings.ToLower("Version")
}
