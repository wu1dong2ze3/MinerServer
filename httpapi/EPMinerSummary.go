package httpapi

import (
	"example.com/m/v2/cgminer"
	"example.com/m/v2/cgminer/body"
	"example.com/m/v2/errs"
	"example.com/m/v2/utils"
	"fmt"
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
	Hashrate   string `json:"hashrate"`   //"hashrate": "473.45 MH/s"
	RunTime    string `json:"RunTime"`    //""duration": 1341242434, //运行时常
	RejectRate string `json:"rejectRate"` //""reject_rate": 0.12, //拒绝率 百分比
	Temp       string `json:"temp"`       //""temp": "58C"//温度，暂时统一成摄氏度
}

func (MinerSummary) GetType() int {
	return GET
}

func (MinerSummary) GetSubPath() string {
	return "/miner/summary"
}

func (MinerSummary) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var result *cgminer.Result
		var errCode *errs.CodeError
		var lang = c.GetHeader("lang")
		var isCn = true
		if result, errCode = cgminer.R(body.Summary{}.ApiCmd(), ""); errCode != nil {
			c.JSON(http.StatusOK, *BaseError(cgminer.CGMinerError))
			return
		}
		var bodys = make([]body.Summary, 0)
		if errCode = result.BindBody(&bodys); errCode != nil {
			c.JSON(http.StatusOK, *BaseError(cgminer.CGMinerError.AddByString("BindBody error!")))
			return
		}
		if len(bodys) != 1 {
			c.JSON(http.StatusOK, *BaseError(cgminer.CGMinerError.AddByString("MinerSummary len(bodys)!=1 error!")))
			return
		}
		if lang == "en-us" {
			isCn = false
		}

		var hashrate = fmt.Sprintf("%.2fGH/s", bodys[0].MHS15m)
		var timeElapsed = ""
		if isCn == true {
			timeElapsed = fmt.Sprintf("%d天 %d小时 %d分 %d秒", utils.DayBySecondRound(bodys[0].Elapsed), utils.HourBySecondMod(bodys[0].Elapsed), utils.MinBySecondMod(bodys[0].Elapsed), utils.SecondBySeccondMod(bodys[0].Elapsed))
		} else {
			timeElapsed = fmt.Sprintf("%d days,%d hours,%d minutes and %d seconds", utils.DayBySecondRound(bodys[0].Elapsed), utils.HourBySecondMod(bodys[0].Elapsed), utils.MinBySecondMod(bodys[0].Elapsed), utils.SecondBySeccondMod(bodys[0].Elapsed))
		}

		var rejectRate = fmt.Sprintf("%.2f%%", bodys[0].DeviceRejected)
		var timeStr = ""
		if timeFloat, errCode := getTemp(); errCode != nil {
		} else {
			timeStr = fmt.Sprintf("%.2fC", timeFloat)
		}

		c.JSON(http.StatusOK, MinerSummary{BaseJson{http.StatusOK, ""}, MinerSummaryResult{SummarySubResult{hashrate, timeElapsed, rejectRate, timeStr}}})
	}
}
func getTemp() (float64, *errs.CodeError) {
	//log.Println("3",cgminer.GetMt().GetCache().Load(time.Duration(10)*time.Second))
	var res, errCod = cgminer.GetMt().GetCache().Load(time.Duration(10) * time.Second)
	if errCod != nil {
		return 0, errCod
	}
	listRes := res.(*[]body.Asc)
	var tempAvg float64
	for _, v := range *listRes {
		tempAvg += v.Temperature
	}
	tempAvg = tempAvg / float64(len(*listRes))
	return tempAvg, nil
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
