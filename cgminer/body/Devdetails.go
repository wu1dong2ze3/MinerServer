package body

import "strings"

type Devdetails struct {
	Devdetails int    `json:"DEVDETAILS"`
	Name       string `json:"Name"`
	ID         int    `json:"ID"`
	Driver     string `json:"Driver"`
	Kernel     string `json:"Kernel"`
	Model      string `json:"Model"`
	DevicePath string `json:"Device Path"`
}

func (v Devdetails) Check() bool {
	return true
}
func (v Devdetails) ApiCmd() string {
	return strings.ToLower("Devdetails")
}
