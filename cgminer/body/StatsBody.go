package body

import "strings"

type Stats struct {
	Stats           int     `json:"STATS"`
	ID              string  `json:"ID"`
	Elapsed         int     `json:"Elapsed"`
	Calls           int     `json:"Calls"`
	Wait            float64 `json:"Wait"`
	Max             float64 `json:"Max"`
	Min             float64 `json:"Min"`
	PoolCalls       int     `json:"Pool Calls"`
	PoolAttempts    int     `json:"Pool Attempts"`
	PoolWait        float64 `json:"Pool Wait"`
	PoolMax         float64 `json:"Pool Max"`
	PoolMin         float64 `json:"Pool Min"`
	PoolAv          float64 `json:"Pool Av"`
	WorkHadRollTime bool    `json:"Work Had Roll Time"`
	WorkCanRoll     bool    `json:"Work Can Roll"`
	WorkHadExpire   bool    `json:"Work Had Expire"`
	WorkRollTime    int     `json:"Work Roll Time"`
	WorkDiff        float64 `json:"Work Diff"`
	MinDiff         float64 `json:"Min Diff"`
	MaxDiff         float64 `json:"Max Diff"`
	MinDiffCount    int     `json:"Min Diff Count"`
	MaxDiffCount    int     `json:"Max Diff Count"`
}

func (v Stats) Check() bool {
	return true
}
func (v Stats) ApiCmd() string {
	return strings.ToLower("Stats")
}
