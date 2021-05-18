package cgminer

import (
	"example.com/m/v2/errs"
	"example.com/m/v2/utils"
	"log"
	"sync"
)

type PoolsJson struct {
	Pools     []PoolJson `json:"pools"`
	APIListen bool       `json:"api-listen"`
	APIPort   string     `json:"api-port"`
	APIAllow  string     `json:"api-allow"`
	T1Pll1    string     `json:"T1Pll1"`
	T1Volt1   string     `json:"T1Volt1"`
	T1Pll2    string     `json:"T1Pll2"`
	T1Volt2   string     `json:"T1Volt2"`
	T1Pll3    string     `json:"T1Pll3"`
	T1Volt3   string     `json:"T1Volt3"`
	T1Pll4    string     `json:"T1Pll4"`
	T1Volt4   string     `json:"T1Volt4"`
}
type PoolJson struct {
	URL  string `json:"url"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

const JsonFile = "/config/cgminer.conf"

var Default = PoolsJson{[]PoolJson{
	{"stratum+tcp://btc.ss.poolin.com:443", "vansbtc.001", "123456"},
	{"stratum+tcp://btc.ss.poolin.com:25", "vansbtc.001", "123456"},
	{"stratum+tcp://btc.ss.poolin.com:1883", "vansbtc.001", "123456"},
}, true, "4028", "W:0/0,W:127.0.0.1",
	"620", "347", "620", "347", "620", "347", "620", "347",
}

type PJM struct{}

var instancePJM *PJM
var oncePJM sync.Once

func InstancePJM() *PJM {
	log.Println("InstanceEPM")
	oncePJM.Do(func() {
		instancePJM = &PJM{}
	})
	return instancePJM
}

func (pjm PJM) Default() *PoolsJson {
	return &Default
}

/**
目前值支持保存用户信息
*/
func (pjm PJM) Save(psJson *PoolsJson) *errs.CodeError {
	if len(psJson.Pools) <= 0 {
		return errs.ParamsError.AddByString("psJson.Pools size at least one!")
	}
	var lastJson *PoolsJson
	var errCode *errs.CodeError
	if lastJson, errCode = pjm.Load(); errCode != nil {
		return errCode
	}
	for _, v := range psJson.Pools {
		if !RegexpPoolUrl.MatchString(v.URL) {
			return errs.ParamsError.AddByString("Url MatchString error!url is " + v.URL)
		}
		if !RegexpUser.MatchString(v.User) || len(v.User) > 255 {
			return errs.ParamsError.AddByString("User is  error!User is " + v.User)
		}
		if v.Pass == "" || len(v.Pass) > 32 {
			return errs.ParamsError.AddByString("Pass is  error!Pass is " + v.Pass)
		}
	}
	copy(lastJson.Pools, psJson.Pools)
	return utils.SaveJson(lastJson, JsonFile)
}

func (pjm PJM) Load() (*PoolsJson, *errs.CodeError) {
	if !utils.IsExist(JsonFile) {
		if errCode := utils.SaveJson(&Default, JsonFile); errCode != nil {
			return nil, errCode
		}
	}
	var pjs = PoolsJson{}
	if errCode := utils.LoadJson(&pjs, JsonFile); errCode != nil {
		return nil, errCode
	}
	return &pjs, nil
}
