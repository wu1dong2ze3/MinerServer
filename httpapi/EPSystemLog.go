package httpapi

import (
	"context"
	"encoding/json"
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
	Token string `json:"token"`
}

type SystemSocketData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Type string `json:"type"`
	Data string `json:"data"`
}

func (SystemLog) GetType() int {
	return GET
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
		if websocket.IsWebSocketUpgrade(c.Request) {
			ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
			if err != nil {
				log.Println(err)
				c.JSON(CreateWebSocketError.Code(), SystemLog{*BaseError(CreateWebSocketError)})
				return
			}
			timeStr := time.Now().Format("2006-01-02 15:04:05")
			log.Println("/system/log=ExecCallBackNoEOF", timeStr)

			cancel := shell.Log.AsynExecCallBack(func(out string, err error, cancel context.CancelFunc) (needContinue bool) {
				log.Println("AsynExecCallBack:", out)
				bytes, _ := json.Marshal(&SystemSocketData{200, "", "normal", out})
				if err = ws.WriteMessage(websocket.TextMessage, bytes); err != nil {
					log.Println("/system/log=AsynExecCallBack error!Over!", err)
					return false
				}
				if err = ws.SetReadDeadline(*utils.AddBaseOnNow(10)); err != nil {
					log.Println("doReadSocket end! SetReadDeadline", err)
					return false
				}
				return true
			})
			go doReadSocket(ws, cancel)
		} else {
			//如果不是websocke默认返回错误
			c.JSON(WebSocketError.Code(), SystemLog{*BaseError(WebSocketError)})
		}
	}
}
func doWriteSocket(ws *websocket.Conn) {

}

func doReadSocket(ws *websocket.Conn, cancel context.CancelFunc) {
	defer ws.Close()
	defer cancel()
	post := &SystemLogPost{}
	if err := ws.SetReadDeadline(*utils.AddBaseOnNow(10)); err != nil {
		log.Println("doReadSocket end!", err)
		cancel()
		return
	}
	for {
		messageType, bytes, err := ws.ReadMessage()
		if err != nil {
			log.Println("doReadSocket end!", err)
			return
		}
		log.Println("doReadSocket any message", string(bytes))
		if err = ws.SetReadDeadline(*utils.AddBaseOnNow(10)); err != nil {
			log.Println("doReadSocket end! SetReadDeadline", err)
			return
		}
		switch messageType {
		case websocket.TextMessage:
			if err = json.Unmarshal(bytes, &post); err != nil {
				log.Println(err)
			}
			break
		case websocket.CloseMessage:
			break
		case websocket.PingMessage:
		case websocket.PongMessage:
			break
		default:
			break

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
