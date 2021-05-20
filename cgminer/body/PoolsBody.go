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
	WorkDifficulty      float64 `json:"Work Difficulty"`
}

func (v Pools) Check() bool {
	return true
}
func (v Pools) ApiCmd() string {
	return strings.ToLower("Pools")
}

/*
{
	"STATUS": [{
		"STATUS": "S",
		"When": 1619621160,
		"Code": 7,
		"Msg": "3 Pool(s)",
		"Description": "cgminer 4.10.0"
	}],
	"POOLS": [{
		"POOL": 0,  //序号
		"URL": "stratum+tcp://btc.ss.poolin.com:443", //矿池地址
		"Status": "Alive", //状态 Alive在线 其他离线
		"Priority": 0,
		"Quota": 1,
		"Long Poll": "N",
		"Getworks": 1023,
		"Accepted": 3, //接收数
		"Rejected": 8, //拒绝数
		"Works": 932553,
		"Discarded": 35215,
		"Stale": 0,
		"Get Failures": 39,
		"Remote Failures": 0,
		"User": "vansbtc.001", //矿共名
		"Last Share Time": 0, //LSTime
		"Diff1 Shares": 0,
		"Proxy Type": "",
		"Proxy": "",
		"Difficulty Accepted": 24576.00000000,
		"Difficulty Rejected": 65536.00000000,
		"Difficulty Stale": 0.00000000,
		"Last Share Difficulty": 0.00000000,
		"Work Difficulty": 65536.00000000, // 难度
		"Has Stratum": true,
		"Stratum Active": true,
		"Stratum URL": "btc.ss.poolin.com",
		"Stratum Difficulty": 65536.00000000,
		"Has Vmask": true,
		"Has GBT": false,
		"Best Share": 0,
		"Pool Rejected%": 72.7273,
		"Pool Stale%": 0.0000,
		"Bad Work": 45,
		"Current Block Height": 680983,
		"Current Block Version": 536870912
	}, {
		"POOL": 1,
		"URL": "stratum+tcp://btc.ss.poolin.com:25",
		"Status": "Alive",
		"Priority": 1,
		"Quota": 1,
		"Long Poll": "N",
		"Getworks": 1,
		"Accepted": 0,
		"Rejected": 0,
		"Works": 1,
		"Discarded": 0,
		"Stale": 0,
		"Get Failures": 0,
		"Remote Failures": 0,
		"User": "vansbtc.002",
		"Last Share Time": 0,
		"Diff1 Shares": 0,
		"Proxy Type": "",
		"Proxy": "",
		"Difficulty Accepted": 0.00000000,
		"Difficulty Rejected": 0.00000000,
		"Difficulty Stale": 0.00000000,
		"Last Share Difficulty": 0.00000000,
		"Work Difficulty": 65536.00000000,
		"Has Stratum": true,
		"Stratum Active": false,
		"Stratum URL": "",
		"Stratum Difficulty": 0.00000000,
		"Has Vmask": true,
		"Has GBT": false,
		"Best Share": 0,
		"Pool Rejected%": 0.0000,
		"Pool Stale%": 0.0000,
		"Bad Work": 0,
		"Current Block Height": 680942,
		"Current Block Version": 536870912
	}, {
		"POOL": 2,
		"URL": "stratum+tcp://btc.ss.poolin.com:1883",
		"Status": "Alive",
		"Priority": 2,
		"Quota": 1,
		"Long Poll": "N",
		"Getworks": 1,
		"Accepted": 0,
		"Rejected": 0,
		"Works": 0,
		"Discarded": 0,
		"Stale": 0,
		"Get Failures": 0,
		"Remote Failures": 0,
		"User": "vansbtc.003",
		"Last Share Time": 0,
		"Diff1 Shares": 0,
		"Proxy Type": "",
		"Proxy": "",
		"Difficulty Accepted": 0.00000000,
		"Difficulty Rejected": 0.00000000,
		"Difficulty Stale": 0.00000000,
		"Last Share Difficulty": 0.00000000,
		"Work Difficulty": 65536.00000000,
		"Has Stratum": true,
		"Stratum Active": false,
		"Stratum URL": "",
		"Stratum Difficulty": 0.00000000,
		"Has Vmask": true,
		"Has GBT": false,
		"Best Share": 0,
		"Pool Rejected%": 0.0000,
		"Pool Stale%": 0.0000,
		"Bad Work": 0,
		"Current Block Height": 680942,
		"Current Block Version": 536870912
	}],
	"id": 1
}
*/
