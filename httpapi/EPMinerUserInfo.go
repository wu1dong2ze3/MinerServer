package httpapi

import (
	"example.com/m/v2/cgminer"
	"example.com/m/v2/errs"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MinerUserInfo struct {
	BaseJson
	Data MinerUserInfoResult `json:"data"`
}

type MinerUserInfoResult struct {
	Miners []UserInfoSubResult `json:"miners"`
}

type UserInfoSubResult struct {
	Index   int    `json:"index"`
	Address string `json:"address"`
	Name    string `json:"name"`
	Pwd     string `json:"pwd"`
}

func (MinerUserInfo) GetType() int {
	return GET
}

func (MinerUserInfo) GetSubPath() string {
	return "/miner/user/info"
}

func (MinerUserInfo) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pjson *cgminer.PoolsJson
		var errCode *errs.CodeError
		if pjson, errCode = cgminer.InstancePJM().Load(); errCode != nil {
			c.JSON(http.StatusOK, *BaseError(errCode))
			return
		}
		mui := MinerUserInfo{*BaseError(errs.NoError), MinerUserInfoResult{[]UserInfoSubResult{}}}
		for index, v := range pjson.Pools {
			mui.Data.Miners = append(mui.Data.Miners, UserInfoSubResult{index, v.URL, v.User, v.Pass})
		}
		c.JSON(http.StatusOK, mui)
	}
}

/**
{
	"code": 200,
	"message": "",
	"data": {
		"miners": [{
			"index": 0,
			"address": "stratum+tcp://btc.ss.poolin.com:443",
			"name": "vansbtc.001",
			"pwd": "123456"
		}, {
			"index": 1,
			"address": "stratum+tcp://btc.ss.poolin.com:25",
			"name": "vansbtc.002",
			"pwd": "123456"
		}, {
			"index": 2,
			"address": "stratum+tcp://btc.ss.poolin.com:1884",
			"name": "vansbtc.003",
			"pwd": "123456"
		}]
	}
}
*/
