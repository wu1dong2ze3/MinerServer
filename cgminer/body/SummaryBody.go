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

/*
{
	"STATUS": [{
		"STATUS": "S",
		"When": 1619619561,
		"Code": 11,
		"Msg": "Summary",
		"Description": "cgminer 4.10.0"
	}],
	"SUMMARY": [{
		"Elapsed": 21761, //TODO 主面板 开机时间？
		"MHS av": 171417.93, //TODO 主面板 实时算力？
		"MHS 5s": 169427.60,
		"MHS 1m": 169147.79,
		"MHS 5m": 168574.95,
		"MHS 15m": 168175.38,
		"Found Blocks": 0,
		"Getworks": 962,
		"Accepted": 3,
		"Rejected": 8,
		"Hardware Errors": 0,
		"Utility": 0.01,
		"Discarded": 32807,
		"Stale": 0,
		"Get Failures": 36,
		"Local Work": 901579,
		"Remote Failures": 0,
		"Network Blocks": 42,
		"Total MH": 3730288588.0000, //TODO 主面板 实时算力？
		"Work Utility": 0.00,
		"Difficulty Accepted": 24576.00000000,
		"Difficulty Rejected": 65536.00000000,
		"Difficulty Stale": 0.00000000,
		"Best Share": 0,
		"Device Hardware%": 0.0000,
		"Device Rejected%": 0.0000, // 主面板 拒绝率
		"Pool Rejected%": 72.7273,
		"Pool Stale%": 0.0000,
		"Last getwork": 1619619561//TODO 主面板 开机时间？
	}],
	"id": 1
}
*/
