package httpapi

import (
	"example.com/m/v2/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserUpdate struct {
	BaseJson
	Data UserUpdateResult `json:"data"`
}
type UserUpdatePost struct {
	User string `json:"user"`
	Pwd  string `json:"pwd"`
	Old  string `json:"old_pwd"`
}

type UserUpdateResult struct {
	Token string `json:"token"`
}

func (UserUpdate) GetType() int {
	return POST
}

func (UserUpdate) GetSubPath() string {
	return "/user/update"
}

func (UserUpdate) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		poster := &UserUpdatePost{}
		var token string
		var err error
		if err = c.ShouldBindJSON(poster); err != nil {
			//c.JSON(http.StatusOK, BaseError(PostJsonError))
			poster.User = c.GetString("user")
			poster.Pwd = c.GetString("pwd")
			poster.Old = c.GetString("old_pwd")
		}
		if err = database.UpdateUser(poster.User, poster.Pwd, poster.Old); err != nil {
			c.JSON(http.StatusForbidden, BaseError(database.UserPwdError))
			return
		}
		if token, err = InstanceTKM().Save(poster.User, poster.Pwd); err != nil {
			c.JSON(http.StatusForbidden, BaseError(database.UserPwdError))
			return
		}
		c.JSON(http.StatusOK, UserUpdate{*BaseError(NoError), UserUpdateResult{token}})
	}
}
