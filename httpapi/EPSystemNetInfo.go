package httpapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

///system/net/info

type SystemNetInfo struct {
	BaseJson
	Data SystemNetInfoResult `json:"data"`
}

type SystemNetInfoResult struct {
	NetInfo NetInfoResult `json:"net_info"`
}

type NetInfoResult struct {
	Mac         string `json:"mac"`         //"mac": "38-AD-45-7E-6D-3E",
	Router_type int    `json:"router_type"` //"router_type": 1,
	Ip          string `json:"ip"`          //"ip": "10.0.0.2",
	SubnetMask  string `json:"subnet_mask"` //"subnet_mask": "255.255.255.0",
	Gateway     string `json:"gateway"`     //"gateway": "10.0.0.1",
	Dns1        string `json:"dns1"`        //"dns1": "114.114.114.114",
	Dns2        string `json:"dns2"`        //"dns2": "202.106.0.20"
}

func (SystemNetInfo) GetType() int {
	return GET
}

func (SystemNetInfo) GetSubPath() string {
	return "/system/net/info"
}

//TODO 假数据
func (SystemNetInfo) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, SystemNetInfo{BaseJson{http.StatusOK, ""},
			SystemNetInfoResult{
				NetInfoResult{"38-AD-45-7E-6D-3E", 1, "10.0.0.2", "255.255.255.0", "10.0.0.1", "114.114.114.114", "202.106.0.20"}}})
	}
}
