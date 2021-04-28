package httpapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

///miner/fan/info
type MinerFanInfo struct {
	BaseJson
	Data MinerFanInfoResult `json:"data"`
}

type MinerFanInfoResult struct {
	Fans []FanInfoSubResult `json:"fans"`
}

type FanInfoSubResult struct {
	Index  int `json:"index"`
	Status int `json:"status"`
	FanRpm int `json:"fanRpm"`
}

func (MinerFanInfo) GetType() int {
	return GET
}

func (MinerFanInfo) GetSubPath() string {
	return "/miner/fan/info"
}

//TODO 假数据
func (MinerFanInfo) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, MinerFanInfo{*BaseError(NoError), MinerFanInfoResult{
			[]FanInfoSubResult{
				{0, 0, 7223},
				{1, 0, 6233},
				{2, 1, 0},
				{3, 0, 5333},
			}},
		})
	}
}

/**
{
  "code": 200,
  "message": "成功",
  "data": {
    "fans": [
      {
        "index": 0,
        "status": 0,
        "fan_rpm": 7233
      },
      {
        "index": 1,
        "status": 0,
        "fan_rpm": 6233
      },
      {
        "index": 2,
        "status": 1,
        "fan_rpm": 0
      },
      {
        "index": 3,
        "status": 0,
        "fan_rpm": 5333
      }
    ]
  }
}
*/
