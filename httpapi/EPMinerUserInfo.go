package httpapi

import (
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

//TODO 假数据
func (MinerUserInfo) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, MinerUserInfo{*BaseError(NoError),
			MinerUserInfoResult{[]UserInfoSubResult{
				{0, "address\": \"stratum+tcp://btc.ss.poolin.com:443", "zhiyuan.5x36", ""},
				{1, "address\": \"stratum+tcp://btc.ss.poolin.com:443", "zhiyuan.5x37", ""},
				{2, "address\": \"stratum+tcp://btc.ss.poolin.com:443", "zhiyuan.5x38", ""},
			}}})
	}
}

/**
{
  "code": 200,
  "message": "成功",
  "data": {
    "miners": [
      {
        "index": 0,
        "address": "stratum+tcp://btc.ss.poolin.com:443",
        "name": "zhiyuan.5x36",
        "pwd": ""
      },
      {
        "index": 1,
        "address": "stratum+tcp://btc.ss.poolin.com:443",
        "name": "zhiyuan.5x37",
        "pwd": ""
      },
      {
        "index": 2,
        "address": "stratum+tcp://btc.ss.poolin.com:443",
        "name": "zhiyuan.5x38",
        "pwd": ""
      }
    ]
  }
}
*/
