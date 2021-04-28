package httpapi

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

///miner/mode/update
type MinerModeUpdate struct {
	BaseJson
}
type MinerModeUpdatePost struct {
	Frequency int     `json:"frequency"`
	Voltage   float64 `json:"voltage"`
}

type MinerModeUpdateResult struct {
}

func (MinerModeUpdate) GetType() int {
	return POST
}

func (MinerModeUpdate) GetSubPath() string {
	return "/miner/mode/update"
}

//TODO 假数据
func (MinerModeUpdate) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		post := MinerModeUpdatePost{}
		var frequency int
		var voltage float64
		if err := c.ShouldBindJSON(&post); err != nil {
			frequency = c.GetInt("frequency")
			voltage = c.GetFloat64("voltage")
		} else {
			frequency = post.Frequency
			voltage = post.Voltage
		}
		log.Println("MinerModeUpdate", frequency, voltage)
		c.JSON(http.StatusOK, MinerModeUpdate{*BaseError(NoError)})
	}
}
