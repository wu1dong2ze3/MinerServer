package body

import "strings"

type Asc struct {
	Asc                 int     `json:"ASC"` // 算力板序号
	Name                string  `json:"Name"`
	ID                  int     `json:"ID"`
	Enabled             string  `json:"Enabled"`
	Status              string  `json:"Status"`      //算力版状态
	Temperature         float64 `json:"Temperature"` //温度
	MhsAv               float64 `json:"MHS av"`
	Mhs5S               float64 `json:"MHS 5s"`
	Mhs1M               float64 `json:"MHS 1m"`
	Mhs5M               float64 `json:"MHS 5m"`
	Mhs15M              float64 `json:"MHS 15m"`
	Accepted            int     `json:"Accepted"`        //接收数
	Rejected            int     `json:"Rejected"`        //拒绝数
	HardwareErrors      int     `json:"Hardware Errors"` //硬件错误
	Utility             float64 `json:"Utility"`
	LastSharePool       int     `json:"Last Share Pool"`
	LastShareTime       int     `json:"Last Share Time"`
	TotalMh             float64 `json:"Total MH"` //实时算力
	Diff1Work           int     `json:"Diff1 Work"`
	DifficultyAccepted  float64 `json:"Difficulty Accepted"`
	DifficultyRejected  float64 `json:"Difficulty Rejected"`
	LastShareDifficulty float64 `json:"Last Share Difficulty"`
	LastValidWork       int     `json:"Last Valid Work"`
	DeviceHardware      float64 `json:"Device Hardware%"`
	DeviceRejected      float64 `json:"Device Rejected%"`
	DeviceElapsed       int     `json:"Device Elapsed"`
	Fminerled           int     `json:"fminerled"`
}

func (v Asc) Check() bool {
	return true
}

func (v Asc) ApiCmd() string {
	return strings.ToLower("asc")
}

/**
{
	"STATUS": [{
		"STATUS": "S",
		"When": 1619709711,
		"Code": 106,
		"Msg": "ASC0",
		"Description": "cgminer 4.10.0"
	}],
	"ASC": [{
		"ASC": 0,
		"Name": "IT1",
		"ID": 0,  // TODO 序号
		"Enabled": "Y",
		"Status": "Alive", //TODO
		"Temperature": 0.00,//TODO
		"MHS av": 42855.49,
		"MHS 5s": 42855.34,
		"MHS 1m": 42855.32,
		"MHS 5m": 42855.82,
		"MHS 15m": 42856.04, //TODO 实时算力 3秒刷一次 g/h  t/t
		"Accepted": 0, //TODO
		"Rejected": 0, //TODO
		"Hardware Errors": 0,  //TODO 硬件错误
		"Utility": 0.00,
		"Last Share Pool": -1,
		"Last Share Time": 0,
		"Total MH": 4796024593.0000,
		"Diff1 Work": 0,
		"Difficulty Accepted": 0.00000000,
		"Difficulty Rejected": 0.00000000,
		"Last Share Difficulty": 0.00000000,
		"Last Valid Work": 1619597839,
		"Device Hardware%": 0.0000,
		"Device Rejected%": 0.0000,
		"Device Elapsed": 111912,
		"fminerled": 0
         //TODO 理论算力暂留
         //TODO 算力板状态
	}],
	"id": 1
}
*/
