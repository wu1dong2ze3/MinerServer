package body

import "strings"

type Devs struct {
	Gpu                 int     `json:"GPU"`
	Enabled             string  `json:"Enabled"`
	Status              string  `json:"Status"`
	Temperature         float64 `json:"Temperature"`
	FanSpeed            int     `json:"Fan Speed"`
	FanPercent          int     `json:"Fan Percent"`
	GpuClock            int     `json:"GPU Clock"`
	MemoryClock         int     `json:"Memory Clock"`
	GpuVoltage          float64 `json:"GPU Voltage"`
	GpuActivity         int     `json:"GPU Activity"`
	Powertune           int     `json:"Powertune"`
	MhsAv               float64 `json:"MHS av"`
	Mhs5S               float64 `json:"MHS 5s"`
	Accepted            int     `json:"Accepted"`
	Rejected            int     `json:"Rejected"`
	HardwareErrors      int     `json:"Hardware Errors"`
	Utility             float64 `json:"Utility"`
	Intensity           string  `json:"Intensity"`
	LastSharePool       int     `json:"Last Share Pool"`
	LastShareTime       int     `json:"Last Share Time"`
	TotalMh             float64 `json:"Total MH"`
	Diff1Work           int     `json:"Diff1 Work"`
	DifficultyAccepted  float64 `json:"Difficulty Accepted"`
	DifficultyRejected  float64 `json:"Difficulty Rejected"`
	LastShareDifficulty float64 `json:"Last Share Difficulty"`
}

func (v Devs) Check() bool {
	return true
}
func (v Devs) ApiCmd() string {
	return strings.ToLower("Devs")
}
