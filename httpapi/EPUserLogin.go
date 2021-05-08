package httpapi

import (
	"example.com/m/v2/database"
	"example.com/m/v2/errs"
	"github.com/gin-gonic/gin"
	"net/http"
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

//TODO 需要实现5次用户名密码校验失败功能
func (UserLogin) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		poster := &UserLoginPost{}
		var err error
		if err = c.ShouldBindJSON(poster); err != nil {
			//c.JSON(http.StatusOK, BaseError(PostJsonError))
			poster.User = c.GetString("user")
			poster.Pwd = c.GetString("pwd")
			poster.Device = c.GetInt("device")
		}
		if err = database.Verify(poster.User, poster.Pwd); err != nil {
			c.JSON(http.StatusOK, BaseError(database.UserPwdError))
			return
		}
		var token = ""
		if token, err = InstanceTKM().Save(poster.User, poster.Pwd); err != nil {
			c.JSON(http.StatusOK, BaseError(database.UserPwdError))
			return
		}
		c.JSON(http.StatusOK, UserLogin{*BaseError(errs.NoError), UserLoginResult{token}})
	}
}
