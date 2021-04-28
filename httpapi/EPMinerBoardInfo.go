package httpapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

///miner/board/info
type MinerBoardInfo struct {
	BaseJson
	Data MinerBoardInfoResult `json:"data"`
}

type MinerBoardInfoResult struct {
	Boards []BoardInfoSubResult `json:"boards"`
}

type BoardInfoSubResult struct {
	Index            int    `json:"index"`
	HashrateRealTime string `json:"hashrateRealTime"`
	HashrateInTheory string `json:"hashrateInTheory"`
	HardwareErrs     int    `json:"hardwareErrs"`
	ReceiveNum       int    `json:"receiveNum"`
	RejectNum        int    `json:"rejectNum"`
	Temp             string `json:"temp"`
	Status           int    `json:"status"`
}

func (MinerBoardInfo) GetType() int {
	return GET
}

func (MinerBoardInfo) GetSubPath() string {
	return "/miner/board/info"
}

//TODO 假数据
func (MinerBoardInfo) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, MinerBoardInfo{*BaseError(NoError),
			MinerBoardInfoResult{[]BoardInfoSubResult{
				{0, "26.39 TH/s", "27.5 TH/s", 425, 1722, 12, "65.1C", 0},
				{1, "26.39 TH/s", "27.5 TH/s", 425, 1722, 12, "65.1C", 0},
				{2, "26.39 TH/s", "27.5 TH/s", 425, 1722, 12, "65.1C", 0},
			}}})
	}
}

/**
{
  "code": 200,
  "message": "成功",
  "data": {
    "boards": [
      {
        "index": 0,
        "hashrate_real_time": "26.39 TH/s",
        "hashrate_in_theory": "27.5 TH/s",
        "hardware_errs": 425,
        "receive_num": 1722,
        "reject_num": 12,
        "temp": "65.1C",
        "status": 0
      },
      {
        "index": 1,
        "hashrate_real_time": "26.39 TH/s",
        "hashrate_in_theory": "27.5 TH/s",
        "hardware_errs": 425,
        "receive_num": 1722,
        "reject_num": 12,
        "temp": "65.1C",
        "status": 0
      },
      {
        "index": 2,
        "hashrate_real_time": "26.39 TH/s",
        "hashrate_in_theory": "27.5 TH/s",
        "hardware_errs": 425,
        "receive_num": 1722,
        "reject_num": 12,
        "temp": "65.1C",
        "status": 0
      }
    ]
  }
}
*/
