package httpapi

import (
	"example.com/m/v2/cgminer"
	"example.com/m/v2/cgminer/body"
	"example.com/m/v2/errs"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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
		var config *body.Config
		if config, errcode = cgminer.Config(); errcode != nil {
			c.JSON(http.StatusOK, *BaseError(errcode))
			return
		}
		log.Println(config)
		if config.ASCCount <= 0 {
			c.JSON(http.StatusOK, *BaseError(cgminer.CGMinerError.AddByString("asc count=0")))
			return
		}
		var resList *[]body.Asc
		if resList, errcode = getAscInfo(config.ASCCount); errcode != nil {
			c.JSON(http.StatusOK, *BaseError(errcode.AddByString("getAscInfo error!")))
			return
		}

		var subRes = make([]BoardInfoSubResult, config.ASCCount)
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

func getAscInfo(ascCount int) (*[]body.Asc, *errs.CodeError) {
	var result *cgminer.Result
	var errCode *errs.CodeError
	var reslist = make([]body.Asc, ascCount)
	for i := 0; i < ascCount; i++ {
		var oneBody []body.Asc
		if result, errCode = cgminer.R(body.Asc{}.ApiCmd(), strconv.Itoa(i)); errCode != nil {
			return nil, errCode
		}
		if errCode = result.BindBody(&oneBody); errCode != nil {
			return nil, errCode
		}
		if len(oneBody) != 1 {
			return nil, cgminer.CGMinerError.AddByString("getAscInfo len(bodyAes)!=1 is " + strconv.Itoa(len(oneBody)))
		}
		reslist = append(reslist, oneBody[0])
	}
	return &reslist, nil
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
