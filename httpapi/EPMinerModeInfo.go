package httpapi

import (
	"example.com/m/v2/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

const ModeJsonFilePath = "/config/mode_info.json"

//miner/mode/info
type MinerModeInfo struct {
	BaseJson
	Data MinerModeInfoResult `json:"data"`
}

type MinerModeInfoResult struct {
	Mode ModeInfoResult `json:"mode"`
}

type ModeInfoResult struct {
	Frequency       int     `json:"frequency"`
	Voltage         float64 `json:"voltage"`
	DefFrequency    int     `json:"defFrequency"`
	DefVoltage      float64 `json:"defVoltage"`
	FreqLimitPct    int     `json:"freqLimitPct"`
	VoltageLimitPct int     `json:"voltageLimitPct"`
}

func (MinerModeInfo) GetType() int {
	return GET
}

func (MinerModeInfo) GetSubPath() string {
	return "/miner/mode/info"
}

var DefaultMR = ModeInfoResult{780, 11.1, 800, 12.34, 20, 20}

func (MinerModeInfo) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		mr := ModeInfoResult{}
		if err := utils.LoadJson(&mr, ModeJsonFilePath); err != nil {
			if err = utils.SaveJson(&DefaultMR, ModeJsonFilePath); err != nil {
				c.JSON(http.StatusOK, *BaseError(err))
				return
			}
			mr = DefaultMR
		}
		c.JSON(http.StatusOK, MinerModeInfo{BaseJson{http.StatusOK, ""},
			MinerModeInfoResult{
				mr}})
	}
}

/**
{
  "code": 200,
  "message": "成功",
  "data": {
    "mode": {
      "frequency": 800,
      "voltage": 12.34
    }
  }
}
*/
