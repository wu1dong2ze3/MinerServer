package httpapi

import (
	"example.com/m/v2/shell"
	"example.com/m/v2/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

///system/log
type SystemLog struct {
	BaseJson
}
type SystemLogPost struct {
	SinceTime string `json:"sinceTime"` //"router_type": 1,
	UntilTime string `json:"untilTime"` //"ip": "10.0.0.2",
}

func (SystemLog) GetType() int {
	return POST
}

func (SystemLog) GetSubPath() string {
	return "/system/log"
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (SystemLog) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("/system/log=begin")
		post := &SystemLogPost{}
		var si, un string
		if err := c.ShouldBindJSON(post); err != nil {
			si = c.GetHeader("sinceTime")
			un = c.GetHeader("untilTime")
		} else {
			si = post.SinceTime
			un = post.UntilTime
		}
		_, err1 := time.Parse("2006-01-02 15:04:05", si)
		_, err2 := time.Parse("2006-01-02 15:04:05", un)
		if (err1 != nil && si != "") || (err2 != nil && un != "") {
			c.JSON(ParamError.Code(), SystemLog{*BaseError(ParamError)})
			return
		}
		if si != "" {
			si = " --since " + "\"" + si + "\""
		}
		if un != "" {
			un = " --until \"" + un + "\""
		}
		if websocket.IsWebSocketUpgrade(c.Request) {
			ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
			if err != nil {
				log.Println(err)
				c.JSON(CreateWebSocketError.Code(), SystemLog{*BaseError(CreateWebSocketError)})
				return
			}
			defer ws.Close()
			if err = ws.SetReadDeadline(*utils.AddBaseOnNow(60)); err != nil {
				return
			}
			timeStr := time.Now().Format("2006-01-02 15:04:05")
			log.Println("/system/log=ExecCallBackNoEOF", timeStr)
			shell.Log.Params(si + un).ExecCallBack(func(out string, err error) (needContinue bool) {
				log.Println(out)
				if err = ws.WriteMessage(websocket.TextMessage, []byte(out)); err != nil {
					log.Println("/system/log=break")
					return false
				}
				return true
			})
			log.Println("/system/log=WriteMessage end at time", timeStr)
			if err = ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(200, timeStr)); err != nil {
				return
			}
			log.Println("/system/log=return")
		} else {
			//如果不是websocke默认返回错误
			c.JSON(WebSocketError.Code(), SystemLog{*BaseError(WebSocketError)})
		}
	}
}

/**```json
{
  "code": 200,
  "message": "成功",
  "data": {
    "log": {
      //共100行
      "start_line": 0,
      //从0行开始读取
      "lines": 2,
      //2
      "log": "Mem: 170420K used, 63324K free, 0K shrd, 643184K buff, 643228K cached\n CPU:   0% usr  20% sys   0% nic  80% idle   0% io   0% irq   0% sirq"
    }
  }
}
```
*/
