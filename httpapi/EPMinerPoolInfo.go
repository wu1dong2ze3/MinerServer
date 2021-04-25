package httpapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MinerPoolInfo struct {
	BaseJson
	Data MinerPoolInfoResult `json:"data"`
}

type MinerPoolInfoResult struct {
	Pools []PoolInfoSubResult `json:"pools"`
}

type PoolInfoSubResult struct {
	Index       int    `json:"index"`       //"index": 0,
	Address     string `json:"address"`     //"address": "stratum+tcp://btc.ss.poolin.com:443",
	Status      int    `json:"status"`      //"status": 1,
	Name        string `json:"name"`        //"name": "zhiyuan.5x36",
	Diff        string `json:"diff"`        //"diff": "262K",
	Receive_num int    `json:"receive_num"` //"receive_num": 2344,
	Reject_num  int    `json:"reject_num"`  //"reject_num": 1,
	Ls_time     int64  `json:"ls_time"`     //"ls_time": 1663445224
}

func (MinerPoolInfo) GetType() int {
	return GET
}

func (MinerPoolInfo) GetSubPath() string {
	return "/miner/pool/info"
}

//TODO 假数据
func (MinerPoolInfo) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, MinerPoolInfo{*BaseError(NoError),
			MinerPoolInfoResult{[]PoolInfoSubResult{
				{0, "stratum+tcp://btc.ss.poolin.com:443", 1, "zhiyuan.5x36", "262K", 2344, 1, 1663445224},
				{1, "stratum+tcp://btc.ss.poolin.com:442", 0, "zhiyuan.5x33", "263K", 2341, 3, 1663445224},
				{2, "stratum+tcp://btc.ss.poolin.com:441", 1, "zhiyuan.5x31", "261K", 0, 1, 1663445224},
			}}})
	}
}

/**
{
  "code": 200,
  "message": "成功",
  "data": {
    "pools": [
      {
        "index": 0,
        "address": "stratum+tcp://btc.ss.poolin.com:443",
        "status": 1,
        "name": "zhiyuan.5x36",
        "diff": "262K",
        "receive_num": 2344,
        "reject_num": 1,
        "ls_time": 1663445224
      },
      {
        "index": 1,
        "address": "stratum+tcp://btc.ss.poolin.com:443",
        "status": 1,
        "name": "zhiyuan.5x37",
        "diff": "262K",
        "receive_num": 2344,
        "reject_num": 1,
        "ls_time": 1663445224
      },
      {
        "index": 2,
        "address": "stratum+tcp://btc.ss.poolin.com:443",
        "status": 1,
        "name": "zhiyuan.5x38",
        "diff": "262K",
        "receive_num": 2344,
        "reject_num": 1,
        "ls_time": 1663445224
      }
    ]
  }
}
*/
