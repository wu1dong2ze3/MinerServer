package httpapi

import (
	"example.com/m/v2/shell"
	"example.com/m/v2/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

///system/hardware/version
type SystemHardwareVersion struct {
	BaseJson
	Data SystemHardwareVersionResult `json:"data"`
}

type SystemHardwareVersionResult struct {
	TitleBar HardwareVersionResult `json:"version"`
}

type HardwareVersionResult struct {
	Model           string `json:"model"`
	Cpu             string `json:"cpu"`
	FirmwareVersion string `json:"firmwareVersion"`
	FirmwareDate    string `json:"firmwareDate"`
}

func (SystemHardwareVersion) GetType() int {
	return GET
}

func (SystemHardwareVersion) GetSubPath() string {
	return "/system/hardware/version"
}

func (SystemHardwareVersion) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		shell.OS_RELEASE.ExecCallBack(func(out string, err error) (needContinue bool) {
			strings.Contains(out, "")
			return false
		})
		version, timeStr, errCode := utils.GetVersionInfo()
		if errCode != nil {
			c.JSON(http.StatusOK, *BaseError(errCode))
			return
		}
		//TODO 假数据
		c.JSON(http.StatusOK, SystemHardwareVersion{BaseJson{http.StatusOK, ""},
			SystemHardwareVersionResult{
				HardwareVersionResult{"测试数据A10s", "测试数据g19", version, timeStr}}})
	}
}

/**
{
  "code": 200,
  "message": "成功",
  "data": {
    "version": {
      "model": "A10s",
      "cpu": "g19",
      "firmware_version": "a10s_20210108_053449",
      "firmware_date": "2020-11-21 15:21"
    }
  }
}
*/
