package httpapi

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

///system/net/update
type SystemNetUpdate struct {
	BaseJson
	Data SystemNetUpdateResult `json:"data"`
}
type SystemNetUpdatePost struct {
	Router_type int    `json:"router_type"` //"router_type": 1,
	Ip          string `json:"ip"`          //"ip": "10.0.0.2",
	SubnetMask  string `json:"subnet_mask"` //"subnet_mask": "255.255.255.0",
	Gateway     string `json:"gateway"`     //"gateway": "10.0.0.1",
	Dns1        string `json:"dns1"`        //"dns1": "114.114.114.114",
	Dns2        string `json:"dns2"`        //"dns2": "202.106.0.20"
}

type SystemNetUpdateResult struct {
	NetInfo NetUpdateResult `json:"net_info"`
}

type NetUpdateResult struct {
	Mac         string `json:"mac"`         //"mac": "38-AD-45-7E-6D-3E",
	Router_type int    `json:"router_type"` //"router_type": 1,
	Ip          string `json:"ip"`          //"ip": "10.0.0.2",
	SubnetMask  string `json:"subnet_mask"` //"subnet_mask": "255.255.255.0",
	Gateway     string `json:"gateway"`     //"gateway": "10.0.0.1",
	Dns1        string `json:"dns1"`        //"dns1": "114.114.114.114",
	Dns2        string `json:"dns2"`        //"dns2": "202.106.0.20"
}

func (SystemNetUpdate) GetType() int {
	return GET
}

func (SystemNetUpdate) GetSubPath() string {
	return "/system/net/info"
}

//TODO 假数据
func (SystemNetUpdate) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		post := SystemNetUpdatePost{}
		var rType int
		var ip, subnet_mask, gateway, dns1, dns2 string
		if err := c.ShouldBindJSON(&post); err != nil {
			rType = c.GetInt("router_type")
			ip = c.GetString("ip")
			subnet_mask = c.GetString("subnet_mask")
			gateway = c.GetString("gateway")
			dns1 = c.GetString("dns1")
			dns2 = c.GetString("dns2")
		} else {
			rType = post.Router_type
			ip = c.GetString("ip")
			subnet_mask = c.GetString("subnet_mask")
			gateway = c.GetString("gateway")
			dns1 = c.GetString("dns1")
			dns2 = c.GetString("dns2")
		}
		log.Println("SystemNetUpdatePost", rType, ip, subnet_mask, gateway, dns1, dns2)
		c.JSON(http.StatusOK, SystemNetUpdate{BaseJson{http.StatusOK, ""},
			SystemNetUpdateResult{
				NetUpdateResult{"38-AD-45-7E-6D-3E", 1, "10.0.0.2", "255.255.255.0", "10.0.0.1", "114.114.114.114", "202.106.0.20"}}})
	}
}
