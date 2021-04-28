package httpapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type MinerSummary struct {
	BaseJson
	Data MinerSummaryResult `json:"data"`
}

type MinerSummaryResult struct {
	Summary SummarySubResult `json:"summary"`
}

type SummarySubResult struct {
	Hashrate   string  `json:"hashrate"`   //"hashrate": "473.45 MH/s"
	StartTime  int64   `json:"startTime"`  //""duration": 1341242434, //运行时常
	RejectRate float64 `json:"rejectRate"` //""reject_rate": 0.12, //拒绝率 百分比
	Temp       string  `json:"temp"`       //""temp": "58C"//温度，暂时统一成摄氏度
	NowTime    int64   `json:"nowTime"`    //""duration": 1341242434, //运行时常
}

func (MinerSummary) GetType() int {
	return GET
}

func (MinerSummary) GetSubPath() string {
	return "/miner/summary"
}

//TODO 假数据
func (MinerSummary) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, MinerSummary{BaseJson{http.StatusOK, ""}, MinerSummaryResult{SummarySubResult{"473.45 MH/s", time.Now().Unix() - 1234567, 0.11, "58C", time.Now().Unix()}}})
	}
}

// 例子
/**
{
  "code": 200,
  "message": "成功",
  "data": {
    "summary": {
      "hashrate": "473.45 MH/s", //实时算力，带单位
      "duration": 1341242434, //运行时常
      "reject_rate": 0.12, //拒绝率 百分比
      "temp": "58C"//温度，暂时统一成摄氏度
    }
  }
}
*/
