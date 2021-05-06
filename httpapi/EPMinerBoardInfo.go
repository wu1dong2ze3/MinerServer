package httpapi

import (
	"example.com/m/v2/cgminer"
	"example.com/m/v2/cgminer/body"
	"example.com/m/v2/errs"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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

func (MinerBoardInfo) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var errcode *errs.CodeError
		var resList *[]body.Asc
		if resList, errcode = getAscInfo(); errcode != nil {
			c.JSON(http.StatusOK, *BaseError(errcode.AddByString("getAscInfo error!")))
			return
		}
		var subRes = make([]BoardInfoSubResult, len(*resList))
		for _, v := range *resList {
			subRes = append(subRes, BoardInfoSubResult{
				Index:            v.ID + 1,
				HashrateRealTime: "需要确认",
				HashrateInTheory: "需要确认",
				HardwareErrs:     v.HardwareErrors,
				ReceiveNum:       v.Accepted,
				RejectNum:        v.Rejected,
				Temp:             fmt.Sprintf("%.1fC", v.Temperature),
				Status:           cgminer.StatusType(v.Status),
			})
		}
		c.JSON(http.StatusOK, MinerBoardInfo{*BaseError(NoError),
			MinerBoardInfoResult{subRes}})
	}
}

func getAscInfo() (*[]body.Asc, *errs.CodeError) {
	var resList *[]body.Asc
	var inter interface{}
	var errCode *errs.CodeError
	//十秒内缓存
	if inter, errCode = cgminer.GetMt().GetCache().Load(time.Duration(10) * time.Second); errCode != nil {
		return nil, errCode
	}
	resList = inter.(*[]body.Asc)
	return resList, nil
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
