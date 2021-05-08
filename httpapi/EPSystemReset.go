package httpapi

import (
	"example.com/m/v2/database"
	"example.com/m/v2/errs"
	"example.com/m/v2/shell"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SystemReset struct {
	BaseJson
}

func (SystemReset) GetType() int {
	return GET
}

func (SystemReset) GetSubPath() string {
	return "/system/reset"
}

func (SystemReset) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		//清除token
		InstanceTKM().Delete()
		//清除数据库
		database.GetInstance().Delete()
		//网络修改成dhcp
		if err := changeNetStatus(1, "", "", "", "", "", func() {
			c.JSON(http.StatusOK, SystemReset{*BaseError(errs.NoError)})
		}); err != nil {
			c.JSON(http.StatusOK, SystemReset{*BaseError(err)})
		}
		shell.STOP_CGMINER.Exec()
		shell.STOP_CGMINER.Exec()
	}
}
