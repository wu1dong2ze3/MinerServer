package httpapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//miner/mode/info
type MinerModeInfo struct {
	BaseJson
	Data MinerModeInfoResult `json:"data"`
}

type MinerModeInfoResult struct {
	Mode ModeInfoResult `json:"mode"`
}

type ModeInfoResult struct {
	Frequency int     `json:"frequency"`
	Voltage   float64 `json:"voltage"`
}

func (MinerModeInfo) GetType() int {
	return GET
}

func (MinerModeInfo) GetSubPath() string {
	return "/miner/mode/info"
}

//TODO 假数据
func (MinerModeInfo) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, MinerModeInfo{BaseJson{http.StatusOK, ""},
			MinerModeInfoResult{
				ModeInfoResult{800, 12.34}}})
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
