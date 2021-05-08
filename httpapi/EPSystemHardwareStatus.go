package httpapi

import (
	"example.com/m/v2/shell"
	"example.com/m/v2/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	Time         string `json:"time"`
	Cpu          int    `json:"cpu"`
	AvailableRam int    `json:"availableRam"`
	TotalRam     int    `json:"totalRam"`
	AvailableRom int    `json:"availableRom"`
	TotalRom     int    `json:"totalRom"`
}

func (SystemHardwareStatus) GetType() int {
	return GET
}

func (SystemHardwareStatus) GetSubPath() string {
	return "/system/hardware/status"
}

func (SystemHardwareStatus) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var time, memAble, memtotal, cpu, rom string
		if time, err = shell.SYSTEMTIME.Exec(); err != nil {
			c.JSON(http.StatusOK, *BaseError(ExeShellError))
			return
		}
		if memAble, err = shell.MemAvailable.Exec(); err != nil {
			c.JSON(http.StatusOK, *BaseError(ExeShellError))
			return
		}
		if memtotal, err = shell.MemTotal.Exec(); err != nil {
			c.JSON(http.StatusOK, *BaseError(ExeShellError))
			return
		}
		if cpu, err = shell.SYSCPU.Exec(); err != nil {
			c.JSON(http.StatusOK, *BaseError(ExeShellError))
			return
		}
		if rom, err = shell.SYSROM.Exec(); err != nil {
			c.JSON(http.StatusOK, *BaseError(ExeShellError))
			return
		}
		time = utils.S{S: time}.NoBr().S
		memAble = utils.S{S: memAble}.NoSpace().NoBr().NoAny("MemAvailable:").S
		memtotal = utils.S{S: memtotal}.NoSpace().NoBr().NoAny("MemTotal:").S
		cpu = utils.S{S: cpu}.NoAny("%").NoAny(`/bin/bash`).NoBr().NoQuot().S
		rom = utils.S{S: rom}.Replace(" ", "/").NoBr().Flip("/").S
		log.Println("cpu", cpu)
		romStr := strings.Split(rom, "/")
		if len(romStr) != 2 {
			c.JSON(http.StatusOK, *BaseError(ExeShellError.AddByString("len(romStr)!=2")))
			return
		}
		c.JSON(http.StatusOK, SystemHardwareStatus{BaseJson{http.StatusOK, ""},
			SystemHardwareStatusResult{
				HardwareStatusResult{time, 100 - UnifyUnits(cpu),
					UnifyUnits(memAble),
					UnifyUnits(memtotal),
					UnifyUnits(romStr[0]),
					UnifyUnits(romStr[1])}}})
	}
}

func UnifyUnits(s string) int {
	var err error
	var f float64
	var ss = ""
	if strings.Contains(s, "kB") {
		ss = utils.S{S: s}.NoAny("kB").S
		if f, err = strconv.ParseFloat(ss, 4); err != nil {
			log.Println("UnifyUnitsToKB error!")
			return 0
		}
		return int(f)
	} else if strings.Contains(s, "M") {
		ss = utils.S{S: s}.NoAny("M").S
		if f, err = strconv.ParseFloat(ss, 4); err != nil {
			log.Println("UnifyUnitsToKB error!")
			return 0
		}
		return int(f) * 1024
	} else if strings.Contains(s, "G") {
		ss = utils.S{S: s}.NoAny("G").S
		if f, err = strconv.ParseFloat(ss, 4); err != nil {
			log.Println("UnifyUnitsToKB error!")
			return 0
		}
		return int(f) * 1024 * 1024
	} else {
		if f, err = strconv.ParseFloat(s, 4); err != nil {
			log.Println("UnifyUnitsToKB error!")
			return 0
		}
		return int(f)
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
