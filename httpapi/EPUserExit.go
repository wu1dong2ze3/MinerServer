package httpapi

import (
	"example.com/m/v2/errs"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserExit struct {
	BaseJson
}

type UserExitResult struct {
	Token string `json:"token"`
}

func (UserExit) GetType() int {
	return GET
}

func (UserExit) GetSubPath() string {
	return "/user/exit"
}

func (UserExit) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		InstanceTKM().Delete()
		c.JSON(http.StatusOK, UserExit{*BaseError(errs.NoError)})
	}
}
