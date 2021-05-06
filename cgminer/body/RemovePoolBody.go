package body

import "strings"

type RemovePool struct {
}

//removepool
func (v RemovePool) Check() bool {
	return true
}

func (v RemovePool) ApiCmd() string {
	return strings.ToLower("removepool")
}

/*
原始json保存 用于测试
{
	"STATUS": [{
		"STATUS": "S",
		"When": 1619762509,
		"Code": 7,
		"Msg": "3 Pool(s)",
		"Description": "cgminer 4.10.0"
	}],
	"POOLS": [{
		"POOL": 0,
		"URL": "stratum+tcp://btc.ss.poolin.com:443",
		"Status": "Alive",
		"Priority": 0,
		"Quota": 1,
		"Long Poll": "N",
		"Getworks": 7178,
		"Accepted": 28,
		"Rejected": 70,
		"Works": 6575217,
		"Discarded": 243943,
		"Stale": 0,
		"Get Failures": 274,
		"Remote Failures": 0,
		"User": "vansbtc.001",
		"Last Share Time": 0,
		"Diff1 Shares": 0,
		"Proxy Type": "",
		"Proxy": "",
		"Difficulty Accepted": 229376.00000000,
		"Difficulty Rejected": 573440.00000000,
		"Difficulty Stale": 0.00000000,
		"Last Share Difficulty": 0.00000000,
		"Work Difficulty": 8192.00000000,
		"Has Stratum": true,
		"Stratum Active": true,
		"Stratum URL": "btc.ss.poolin.com",
		"Stratum Difficulty": 8192.00000000,
		"Has Vmask": true,
		"Has GBT": false,
		"Best Share": 0,
		"Pool Rejected%": 71.4286,
		"Pool Stale%": 0.0000,
		"Bad Work": 274,
		"Current Block Height": 681191,
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
