package httpapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

///miner/user/update
type MinerUserUpdate struct {
	BaseJson
}

type MinerUserUpdatePost struct {
	Miner []UserUpdateSubPost `json:"miner"`
}

type UserUpdateSubPost struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Pwd     string `json:"pwd"`
}

func (MinerUserUpdate) GetType() int {
	return POST
}

func (MinerUserUpdate) GetSubPath() string {
	return "/miner/user/update"
}

//TODO 假数据
func (MinerUserUpdate) GetHandle() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, MinerUserUpdate{*BaseError(NoError)})
	}
}
