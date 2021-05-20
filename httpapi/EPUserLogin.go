package httpapi

import (
	"example.com/m/v2/database"
	"example.com/m/v2/errs"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserLogin struct {
	BaseJson
	Data UserLoginResult `json:"data"`
}
type UserLoginPost struct {
	User   string `json:"user"`
	Pwd    string `json:"pwd"`
	Device int    `json:"device"`
}

type UserLoginResult struct {
	Token string `json:"token"`
}

func (UserLogin) GetType() int {
	return POST
}

func (UserLogin) GetSubPath() string {
	return "/user/login"
}

func (UserLogin) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIp := c.ClientIP()
		fmt.Println("login:", clientIp)
		poster := &UserLoginPost{}
		var err error
		if err = c.ShouldBindJSON(poster); err != nil {
			poster.User = c.GetString("user")
			poster.Pwd = c.GetString("pwd")
			poster.Device = c.GetInt("device")
		}
		if !oneMistake(c.ClientIP()) {
			c.JSON(http.StatusOK, *BaseError(ErrTooManyRetries.AddByString("Ip is " + clientIp)))
			return
		}
		if err = database.Verify(poster.User, poster.Pwd); err != nil {
			c.JSON(http.StatusOK, *BaseError(database.UserPwdError))
			return
		}
		var token = ""
		if token, err = InstanceTKM().Save(poster.User, poster.Pwd); err != nil {
			c.JSON(http.StatusOK, *BaseError(database.UserPwdError))
			return
		}
		deleteMistake(clientIp)
		c.JSON(http.StatusOK, UserLogin{*BaseError(errs.NoError), UserLoginResult{token}})
	}
}

type client struct {
	wrongTimes int
	time       time.Time
}

var clients = make(map[string]*client)

/**
true 允许通过校验，false 不允许
*/

func oneMistake(ip string) bool {
	cl, ok := clients[ip]
	//第一次记录
	if !ok {
		clients[ip] = &client{1, time.Now()}
		fmt.Println("oneMistake:", ip)
		return true
	}
	//判断记录时间是否超时
	if cl.time.Add(5 * time.Minute).After(time.Now()) {
		//5分钟以内判断失败次数
		cl.wrongTimes = cl.wrongTimes + 1
		if cl.wrongTimes >= 5 {
			fmt.Printf("times %d \n", cl.wrongTimes)
			return false
		} else {
			fmt.Printf("times %d \n", cl.wrongTimes)
			return true
		}

	} else {
		//超过5分钟刷新时间，重新计数
		cl.time = time.Now()
		cl.wrongTimes = 1
		fmt.Println("over time oneMistake:", ip)
		return true
	}
}

func deleteMistake(ip string) {
	delete(clients, ip)
}
