package httpapi

import (
	"example.com/m/v2/shell"
	"example.com/m/v2/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

///system/ui/title_bar
type SystemUiTitleBar struct {
	BaseJson
	Data SystemUiTitleBarResult `json:"data"`
}

type SystemUiTitleBarResult struct {
	TitleBar UiTitleBarResult `json:"titleBar"`
}

type UiTitleBarResult struct {
	MinerName string `json:"minerName"`
	Ip        string `json:"ip"`
	Mac       string `json:"mac"`
	Version   string `json:"version"`
}

func (SystemUiTitleBar) GetType() int {
	return GET
}

func (SystemUiTitleBar) GetSubPath() string {
	return "/system/ui/titleBar"
}

func (SystemUiTitleBar) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		version, _, errCode := utils.GetVersionInfo()
		if errCode != nil {
			c.JSON(http.StatusOK, *BaseError(errCode))
			return
		}
		var ip, mac string
		shell.NetworkStatus.ExecCallBack(func(out string, err error) (needContinue bool) {
			if strings.Contains(out, "HW Address: ") {
				arr := strings.Split(out, ": ")
				log.Println("HW Address:", arr)
				mac = utils.S{S: arr[1]}.NoSpaceBr().S
				return true
			}
			if !strings.Contains(out, "HW Address: ") && strings.Contains(out, "Address: ") {
				arr := strings.Split(out, ": ")
				log.Println("Address:", arr)
				ip = utils.S{S: arr[1]}.NoSpaceBr().S
				return true
			}
			return true
		})
		c.JSON(http.StatusOK, SystemUiTitleBar{BaseJson{http.StatusOK, ""},
			SystemUiTitleBarResult{
				UiTitleBarResult{"PoolinMiner1", ip, mac, version}}})
	}
}

/**

```json
{
  "code": 200,
  "message": "成功",
  "data": {
    "title_bar": {
      "ip": "192.168.1.1",
      "mac": "C4:E5:E7:3F:E0:F4",
      "version": "a10s_20210108_053449"
    }
  }
}
```
*/
