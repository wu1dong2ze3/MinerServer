package body

import "strings"

type Pools struct {
	Pool                int     `json:"POOL"`
	URL                 string  `json:"URL"`
	Status              string  `json:"Status"`
	Priority            int     `json:"Priority"`
	LongPoll            string  `json:"Long Poll"`
	Getworks            int     `json:"Getworks"`
	Accepted            int     `json:"Accepted"`
	Rejected            int     `json:"Rejected"`
	Discarded           int     `json:"Discarded"`
	Stale               int     `json:"Stale"`
	GetFailures         int     `json:"Get Failures"`
	RemoteFailures      int     `json:"Remote Failures"`
	User                string  `json:"User"`
	LastShareTime       int     `json:"Last Share Time"`
	Diff1Shares         int     `json:"Diff1 Shares"`
	ProxyType           string  `json:"Proxy Type"`
	Proxy               string  `json:"Proxy"`
	DifficultyAccepted  float64 `json:"Difficulty Accepted"`
	DifficultyRejected  float64 `json:"Difficulty Rejected"`
	DifficultyStale     float64 `json:"Difficulty Stale"`
	LastShareDifficulty float64 `json:"Last Share Difficulty"`
	HasStratum          bool    `json:"Has Stratum"`
	StratumActive       bool    `json:"Stratum Active"`
	StratumURL          string  `json:"Stratum URL"`
}

func (v Pools) Check() bool {
	return true
}
func (v Pools) ApiCmd() string {
	return strings.ToLower("Pools")
}
