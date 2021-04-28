package body

import "strings"

type Coin struct {
	HashMethod       string  `json:"Hash Method"`
	CurrentBlockTime float64 `json:"Current Block Time"`
	CurrentBlockHash string  `json:"Current Block Hash"`
	Lp               bool    `json:"LP"`
}

func (v Coin) Check() bool {
	return true
}
func (v Coin) ApiCmd() string {
	return strings.ToLower("Coin")
}
