package httpapi

import (
	"example.com/m/v2/errs"
	"example.com/m/v2/shell"
	"example.com/m/v2/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

///system/ota/info

type SystemOteInfo struct {
	BaseJson
	Data SystemOteInfoResult `json:"data"`
}

type SystemOteInfoResult struct {
	Ota OteInfoResult `json:"ota"`
}

type OteInfoResult struct {
	Version string `json:"version"`
	Time    string `json:"time"`
}

func (SystemOteInfo) GetType() int {
	return GET
}

func (SystemOteInfo) GetSubPath() string {
	return "/system/ota/info"
}

const configFile = "/etc/sw_info"

func (SystemOteInfo) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !utils.IsExist(configFile) {
			c.JSON(http.StatusOK, *BaseError(errs.IoError.AddByString(configFile + " is not exist!")))
			return
		}
		var err error
		var ini *utils.Ini
		if ini, err = utils.Open(configFile); err != nil {
			c.JSON(http.StatusOK, *BaseError(errs.IoError.AddByString(configFile + " can't open!")))
			return
		}
		var version string
		if version = ini.Load("sw_version", ""); version == "" {
			c.JSON(http.StatusOK, *BaseError(errs.IoError.AddByString("sw_version is null")))
			return
		}

		var timeStr string
		shell.STAT.Params(configFile).ExecCallBack(func(out string, err error) (needContinue bool) {
			if strings.Contains(out, "Modify:") {
				timeStr = out
				timeStr = utils.S{timeStr}.NoAny("Modify:").NoBr().Select(" ", " ", 1, 2).Select(".", "", 0).S
				return false
			}
			return true
		})

		c.JSON(http.StatusOK, SystemOteInfo{BaseJson{http.StatusOK, ""},
			SystemOteInfoResult{
				OteInfoResult{version, timeStr}}})
	}
}

/**
```json
{
  "code": 200,
  "message": "成功",
  "data": {
    "ota": {
      "version": "ota-aa-bb-cc-1",
      "time": "2021.4.5 22:22",
      "ota_name": "/user/file/ota_01_01.rar"
    }
  }
}
```
*/
