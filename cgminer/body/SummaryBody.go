package body

import "strings"

type Summary struct {
	Elapsed            int     `json:"Elapsed"`
	Algorithm          string  `json:"Algorithm"`
	MhsAv              float64 `json:"MHS av"`
	FoundBlocks        int     `json:"Found Blocks"`
	Getworks           int     `json:"Getworks"`
	Accepted           int     `json:"Accepted"`
	Rejected           int     `json:"Rejected"`
	HardwareErrors     int     `json:"Hardware Errors"`
	Utility            float64 `json:"Utility"`
	Discarded          int     `json:"Discarded"`
	Stale              int     `json:"Stale"`
	GetFailures        int     `json:"Get Failures"`
	LocalWork          int     `json:"Local Work"`
	RemoteFailures     int     `json:"Remote Failures"`
	NetworkBlocks      int     `json:"Network Blocks"`
	TotalMh            float64 `json:"Total MH"`
	WorkUtility        float64 `json:"Work Utility"`
	DifficultyAccepted float64 `json:"Difficulty Accepted"`
	DifficultyRejected float64 `json:"Difficulty Rejected"`
	DifficultyStale    float64 `json:"Difficulty Stale"`
}

func (v Summary) Check() bool {
	return true
}
func (v Summary) ApiCmd() string {
	return strings.ToLower("Summary")
}
