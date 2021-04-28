package httpapi

import (
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

//TODO 假数据
func (SystemReset) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, SystemReset{*BaseError(NoError)})
	}
}
