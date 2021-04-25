package body

import "strings"

type Config struct {
	GpuCount     int    `json:"GPU Count"`
	PgaCount     int    `json:"PGA Count"`
	CPUCount     int    `json:"CPU Count"`
	PoolCount    int    `json:"Pool Count"`
	Adl          string `json:"ADL"`
	AdlInUse     string `json:"ADL in use"`
	Strategy     string `json:"Strategy"`
	LogInterval  int    `json:"Log Interval"`
	DeviceCode   string `json:"Device Code"`
	Os           string `json:"OS"`
	FailoverOnly bool   `json:"Failover-Only"`
	Scantime     int    `json:"ScanTime"`
	Queue        int    `json:"Queue"`
	Expiry       int    `json:"Expiry"`
}

func (v Config) Check() bool {
	return true
}

func (v Config) ApiCmd() string {
	return strings.ToLower("Config")
}
