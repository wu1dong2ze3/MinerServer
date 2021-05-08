package httpapi

import (
	"example.com/m/v2/errs"
	"example.com/m/v2/shell"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SystemReboot struct {
	BaseJson
}

func (SystemReboot) GetType() int {
	return GET
}

func (SystemReboot) GetSubPath() string {
	return "/system/reboot"
}

func (SystemReboot) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, SystemReboot{*BaseError(errs.NoError)})
		if _, err := shell.Reboot.Exec(); err != nil {
			c.JSON(MinerRebootFailed.Code(), SystemReboot{*BaseError(MinerRebootFailed.Add(err))})
			return
		}
	}
}
