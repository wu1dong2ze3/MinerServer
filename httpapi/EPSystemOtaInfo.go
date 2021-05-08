package httpapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

///system/ota/info

type SystemOteInfo struct {
	BaseJson
	Data SystemOteInfoResult `json:"data"`
}

type SystemOteInfoResult struct {
	ota OteInfoResult `json:"ota"`
}

type OteInfoResult struct {
	Version string `json:"version"`
	Time    string `json:"time"`
	OtaName string `json:"otaName"`
}

func (SystemOteInfo) GetType() int {
	return GET
}

func (SystemOteInfo) GetSubPath() string {
	return "/system/ota/info"
}

//TODO 假数据
func (SystemOteInfo) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, SystemOteInfo{BaseJson{http.StatusOK, ""},
			SystemOteInfoResult{
				OteInfoResult{"测试数据ota-aa-bb-cc-1", "测试数据2021.4.5 22:22", "测试数据/user/file/ota_01_01.rar"}}})
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
