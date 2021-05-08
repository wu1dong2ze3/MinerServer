package httpapi

import (
	"example.com/m/v2/database"
	"example.com/m/v2/errs"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserUpdate struct {
	BaseJson
	Data UserUpdateResult `json:"data"`
}
type UserUpdatePost struct {
	Pwd string `json:"newPwd"`
	Old string `json:"oldPwd"`
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
			poster.Pwd = c.GetString("newPwd")
			poster.Old = c.GetString("oldPwd")
		}
		var user *database.User
		if user, err = database.UpdateUser(poster.Pwd, poster.Old); err != nil {
			c.JSON(http.StatusForbidden, BaseError(database.UserPwdError))
			return
		}
		if token, err = InstanceTKM().Save(user.Name, user.Pwd); err != nil {
			c.JSON(http.StatusForbidden, BaseError(database.UserPwdError))
			return
		}
		c.JSON(http.StatusOK, UserUpdate{*BaseError(errs.NoError), UserUpdateResult{token}})
	}
}
