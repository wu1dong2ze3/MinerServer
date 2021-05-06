package httpapi

import (
	"example.com/m/v2/shell"
	"example.com/m/v2/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

///system/hardware/status
type SystemHardwareStatus struct {
	BaseJson
	Data SystemHardwareStatusResult `json:"data"`
}

type SystemHardwareStatusResult struct {
	Status HardwareStatusResult `json:"status"`
}

type HardwareStatusResult struct {
	Time string `json:"time"`
	Cpu  string `json:"cpu"`
	Ram  string `json:"ram"`
	Rom  string `json:"rom"`
}

func (SystemHardwareStatus) GetType() int {
	return GET
}

func (SystemHardwareStatus) GetSubPath() string {
	return "/system/hardware/status"
}

//TODO 需要处理
func (SystemHardwareStatus) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var time, memAble, memtotal, cpu, rom string
		var cpui int
		if time, err = shell.SYSTEMTIME.Exec(); err != nil {
			c.JSON(http.StatusOK, *BaseError(CanNotLoadFile))
			return
		}
		if memAble, err = shell.MemAvailable.Exec(); err != nil {
			c.JSON(http.StatusOK, *BaseError(CanNotLoadFile))
			return
		}
		if memtotal, err = shell.MemTotal.Exec(); err != nil {
			c.JSON(http.StatusOK, *BaseError(CanNotLoadFile))
			return
		}
		if cpu, err = shell.SYSCPU.Exec(); err != nil {
			c.JSON(http.StatusOK, *BaseError(CanNotLoadFile))
			return
		}
		if rom, err = shell.SYSROM.Exec(); err != nil {
			c.JSON(http.StatusOK, *BaseError(CanNotLoadFile))
			return
		}
		time = utils.S{S: time}.NoBr().S
		memAble = utils.S{S: memAble}.NoSpace().NoBr().NoAny("MemAvailable:").S
		memtotal = utils.S{S: memtotal}.NoSpace().NoBr().NoAny("MemTotal:").S
		cpu = utils.S{S: cpu}.NoAny("%").NoAny(`/bin/bash`).NoBr().NoQuot().S
		rom = utils.S{S: rom}.Replace(" ", "/").NoBr().Flip("/").S
		log.Println("cpu", cpu)
		if cpui, err = strconv.Atoi(cpu); err != nil {
			c.JSON(http.StatusOK, *BaseError(CanNotLoadFile.Add(err)))
			return
		}
		cpu = strconv.Itoa(100-cpui) + "%"
		c.JSON(http.StatusOK, SystemHardwareStatus{BaseJson{http.StatusOK, ""},
			SystemHardwareStatusResult{
				HardwareStatusResult{time, cpu, memAble + "/" + memtotal, rom}}})
	}
}

/**
{
  "code": 200,
  "message": "成功",
  "data": {
    "status": {
      "time": "2021 03-29 09:39:51",
      //开机时间
      "cpu": "50",
      //cpu占用 （百分比）
      "ram": "63108/240384",
      //内存占用
      "rom": "34368/223144"
      //存储占用
    }
  }
}
*/
