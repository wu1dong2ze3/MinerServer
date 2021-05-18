package httpapi

import (
	"example.com/m/v2/utils"
	"github.com/gin-gonic/gin"
	"net/http"
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
		version, timeStr, errCode := utils.GetVersionInfo()
		if errCode != nil {
			c.JSON(http.StatusOK, *BaseError(errCode))
			return
		}
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
