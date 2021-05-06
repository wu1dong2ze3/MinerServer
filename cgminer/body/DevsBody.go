package body

import "strings"

type Devs struct {
	ID                  int     `json:"ID"`
	Enabled             string  `json:"Enabled"`
	Status              string  `json:"Status"`
	Temperature         float64 `json:"Temperature"`
	FanSpeed            int     `json:"Fan Speed"` //风扇转数
	FanPercent          int     `json:"Fan Percent"`
	GpuClock            int     `json:"GPU Clock"`
	MemoryClock         int     `json:"Memory Clock"`
	GpuVoltage          float64 `json:"GPU Voltage"`
	GpuActivity         int     `json:"GPU Activity"`
	Powertune           int     `json:"Powertune"`
	MhsAv               float64 `json:"MHS av"`
	Mhs5S               float64 `json:"MHS 5s"`
	MHS15m              float64 `json:"MHS 15m"`
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

/*
{
	"STATUS": [{
		"STATUS": "S",
		"When": 1619621434,
		"Code": 9,
		"Msg": "4 ASC(s)",
		"Description": "cgminer 4.10.0"
	}],
	"DEVS": [{
		"ASC": 0,
		"Name": "IT1",
		"ID": 0, // TODO 序号
		"Enabled": "Y",
		"Status": "Alive", //TODO
		"Temperature": 0.00, //TODO
		"MHS av": 42854.64,
		"MHS 5s": 42856.93,
		"MHS 1m": 42856.46,
		"MHS 5m": 42856.30,
		"MHS 15m": 42856.30, //TODO 实时算力 5秒刷一次 g/h  t/t
		"Accepted": 0, //TODO
		"Rejected": 0, //TODO
		"Hardware Errors": 0, //TODO 硬件错误
		"Utility": 0.00,
		"Last Share Pool": -1,
		"Last Share Time": 0,
		"Total MH": 1012836352.0000,
		"Diff1 Work": 0,
		"Difficulty Accepted": 0.00000000,
		"Difficulty Rejected": 0.00000000,
		"Last Share Difficulty": 0.00000000,
		"Last Valid Work": 1619597804,
		"Device Hardware%": 0.0000,
		"Device Rejected%": 0.0000,
		"Device Elapsed": 23634,
		"fminerled": 0
	}, {
		"ASC": 1,
		"Name": "IT1",
		"ID": 1,
		"Enabled": "Y",
		"Status": "Alive",
		"Temperature": 0.00,
		"MHS av": 42854.64,
		"MHS 5s": 42856.27,
		"MHS 1m": 42856.40,
		"MHS 5m": 42856.29,
		"MHS 15m": 42856.30,
		"Accepted": 0,
		"Rejected": 0,
		"Hardware Errors": 0,
		"Utility": 0.00,
		"Last Share Pool": -1,
		"Last Share Time": 0,
		"Total MH": 1012836352.0000,
		"Diff1 Work": 0,
		"Difficulty Accepted": 0.00000000,
		"Difficulty Rejected": 0.00000000,
		"Last Share Difficulty": 0.00000000,
		"Last Valid Work": 1619597807,
		"Device Hardware%": 0.0000,
		"Device Rejected%": 0.0000,
		"Device Elapsed": 23634,
		"fminerled": 0
	}, {
		"ASC": 2,
		"Name": "IT1",
		"ID": 2,
		"Enabled": "Y",
		"Status": "Alive",
		"Temperature": 0.00,
		"MHS av": 42854.64,
		"MHS 5s": 42857.38,
		"MHS 1m": 42856.54,
		"MHS 5m": 42856.32,
		"MHS 15m": 42856.31,
		"Accepted": 0,
		"Rejected": 0,
		"Hardware Errors": 0,
		"Utility": 0.00,
		"Last Share Pool": -1,
		"Last Share Time": 0,
		"Total MH": 1012836352.0000,
		"Diff1 Work": 0,
		"Difficulty Accepted": 0.00000000,
		"Difficulty Rejected": 0.00000000,
		"Last Share Difficulty": 0.00000000,
		"Last Valid Work": 1619597810,
		"Device Hardware%": 0.0000,
		"Device Rejected%": 0.0000,
		"Device Elapsed": 23634,
		"fminerled": 0
	}, {
		"ASC": 3,
		"Name": "IT1",
		"ID": 3,
		"Enabled": "Y",
		"Status": "Alive",
		"Temperature": 0.00,
		"MHS av": 42854.64,
		"MHS 5s": 42856.90,
		"MHS 1m": 42856.45,
		"MHS 5m": 42856.29,
		"MHS 15m": 42856.30,
		"Accepted": 0,
		"Rejected": 0,
		"Hardware Errors": 0,
		"Utility": 0.00,
		"Last Share Pool": -1,
		"Last Share Time": 0,
		"Total MH": 1012836352.0000,
		"Diff1 Work": 0,
		"Difficulty Accepted": 0.00000000,
		"Difficulty Rejected": 0.00000000,
		"Last Share Difficulty": 0.00000000,
		"Last Valid Work": 1619597809,
		"Device Hardware%": 0.0000,
		"Device Rejected%": 0.0000,
		"Device Elapsed": 23634,
		"fminerled": 0
	}],
	"id": 1
}
*/
