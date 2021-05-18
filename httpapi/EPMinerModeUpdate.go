package httpapi

import (
	"example.com/m/v2/cgminer"
	"example.com/m/v2/cgminer/body"
	"example.com/m/v2/errs"
	"example.com/m/v2/shell"
	"example.com/m/v2/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
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

func (MinerModeUpdate) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		post := &MinerModeUpdatePost{}
		var frequency int
		var voltage float64
		if err := c.ShouldBindJSON(post); err != nil {
			frequency = c.GetInt("frequency")
			voltage = c.GetFloat64("voltage")
		} else {
			frequency = post.Frequency
			voltage = post.Voltage
		}
		mr := ModeInfoResult{}
		if err := utils.LoadJson(&mr, ModeJsonFilePath); err != nil {
			c.JSON(http.StatusOK, *BaseError(err))
			return
		}

		if err := check(frequency, voltage, mr); err != nil {
			c.JSON(http.StatusOK, *BaseError(err))
			return
		}
		mr.Frequency = frequency
		mr.Voltage = voltage
		if err := utils.SaveJson(&mr, ModeJsonFilePath); err != nil {
			c.JSON(http.StatusOK, *BaseError(err))
			return
		}
		if _, err := shell.RESTART_CGMINER.Exec(); err != nil {
			c.JSON(http.StatusOK, *BaseError(CgMinerDeviceError.Add(err)))
			return
		} else {
			testOk := false
			for i := 0; i < 10; i++ {
				if _, errCode := cgminer.R(body.Version{}.ApiCmd(), ""); errCode != nil {
					time.Sleep(time.Duration(2) * time.Second)
					log.Printf("try to restart! %d times,error is %s", i+1, errCode.Error())
					continue
				}
				testOk = true
				break
			}
			if testOk {
				c.JSON(http.StatusOK, *BaseError(errs.NoError))
			} else {
				c.JSON(http.StatusOK, *BaseError(CgMinerDeviceError.AddByString("cgminer restart over runtime!")))
			}
			return
		}
	}
}

func check(freg int, voltage float64, mr ModeInfoResult) *errs.CodeError {
	downfreg := int(float64(mr.DefFrequency) - float64(mr.DefFrequency)*(float64(mr.FreqLimitPct)/float64(100)))
	upfreg := int(float64(mr.DefFrequency) + float64(mr.DefFrequency)*(float64(mr.FreqLimitPct)/float64(100)))
	if freg < downfreg || freg > upfreg {
		return UpdateFregError.AddByString(fmt.Sprintf("freg %d is error,Out of range!", freg))
	}
	downV := float64(mr.DefVoltage) - float64(mr.DefVoltage)*(float64(mr.VoltageLimitPct)/float64(100))
	upV := float64(mr.DefVoltage) + float64(mr.DefVoltage)*(float64(mr.VoltageLimitPct)/float64(100))
	if voltage < downV || voltage > upV {
		return UpdateVolError.AddByString(fmt.Sprintf("voltage %d is error,Out of range!", voltage))
	}

	return nil
}
