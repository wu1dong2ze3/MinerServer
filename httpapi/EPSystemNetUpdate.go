package httpapi

import (
	"example.com/m/v2/shell"
	"example.com/m/v2/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
)

///system/net/update
type SystemNetUpdate struct {
	BaseJson
}
type SystemNetUpdatePost struct {
	RouterType int    `json:"routerType"` //"router_type": 1,
	Ip         string `json:"ip"`         //"ip": "10.0.0.2",
	SubnetMask string `json:"subnetMask"` //"subnet_mask": "255.255.255.0",
	Gateway    string `json:"gateway"`    //"gateway": "10.0.0.1",
	Dns1       string `json:"dns1"`       //"dns1": "114.114.114.114",
	Dns2       string `json:"dns2"`       //"dns2": "202.106.0.20"
}

func (SystemNetUpdate) GetType() int {
	return POST
}

func (SystemNetUpdate) GetSubPath() string {
	return "/system/net/update"
}

func (SystemNetUpdate) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("SystemNetUpdatePost a")
		post := SystemNetUpdatePost{}
		var rType int
		var ip, subnet_mask, gateway, dns1, dns2 string
		var err error
		if err = c.ShouldBindJSON(&post); err != nil {
			rType = c.GetInt("routerType")
			ip = c.GetString("ip")
			subnet_mask = c.GetString("subnetMask")
			gateway = c.GetString("gateway")
			dns1 = c.GetString("dns1")
			dns2 = c.GetString("dns2")
		} else {
			rType = post.RouterType
			ip = c.GetString("ip")
			subnet_mask = c.GetString("subnetMask")
			gateway = c.GetString("gateway")
			dns1 = c.GetString("dns1")
			dns2 = c.GetString("dns2")
		}
		var ipcdr string
		if rType == 0 {
			if ipcdr, err = utils.CheckCidr(ip, subnet_mask); err != nil {
				c.JSON(http.StatusOK, SystemNetUpdate{*BaseError(utils.ParamError)})
				return
			}
			ip = ipcdr
			if gateway != "" {
				if ip := net.ParseIP(gateway); ip == nil {
					c.JSON(http.StatusOK, SystemNetUpdate{*BaseError(utils.ParamError)})
					return
				}
			}
			if dns1 != "" {
				if ip := net.ParseIP(dns1); ip == nil {
					c.JSON(http.StatusOK, SystemNetUpdate{*BaseError(utils.ParamError)})
					return
				}
			}
			if dns2 != "" {
				if ip := net.ParseIP(dns2); ip == nil {
					c.JSON(http.StatusOK, SystemNetUpdate{*BaseError(utils.ParamError)})
					return
				}
			}
		}
		log.Println("SystemNetUpdatePost", rType, ip, subnet_mask, gateway, dns1, dns2)
		//static
		if rType == 0 {
			ini, err := utils.New("/config/network/25-wired.network")
			if err != nil {
				c.JSON(http.StatusOK, SystemNetUpdate{*BaseError(utils.ParamError)})
				return
			}
			err = ini.Add("[Match]", "")
			err = ini.Add("Name", "eth0")
			err = ini.Add("[Network]", "")
			err = ini.Add("Netmask", subnet_mask)
			err = ini.Add("Address", ipcdr)
			if gateway != "" {
				err = ini.Add("Gateway", gateway)
			}
			if dns1 != "" {
				err = ini.Add("DNS", dns1)
			}
			if dns2 != "" {
				err = ini.Add("DNS", dns2)
			}
			err = ini.Save()
			if err != nil {
				c.JSON(http.StatusOK, SystemNetUpdate{*BaseError(utils.ParamError)})
				return
			}
			_, err = shell.NetworkRestart.Exec()
			if err != nil {
				c.JSON(http.StatusOK, SystemNetUpdate{*BaseError(utils.ParamError)})
				return
			}
			//dhcp
		} else if rType == 1 {
			ini, err := utils.New("/config/network/25-wired.network")
			if err != nil {
				c.JSON(http.StatusOK, SystemNetUpdate{*BaseError(utils.ParamError)})
				return
			}
			err = ini.Add("[Match]", "")
			err = ini.Add("Name", "eth0")
			err = ini.Add("[Network]", "")
			err = ini.Add("DHCP", "yes")
			err = ini.Add("[DHCP]", "")
			err = ini.Add("ClientIdentifier", "mac")
			if err != nil {
				c.JSON(http.StatusOK, SystemNetUpdate{*BaseError(utils.ParamError)})
				return
			}
			c.JSON(http.StatusOK, SystemNetUpdate{BaseJson{http.StatusOK, ""}})
			_, err = shell.NetworkRestart.Exec()
			if err != nil {
				c.JSON(http.StatusOK, SystemNetUpdate{*BaseError(utils.ParamError)})
				return
			}
		}

	}
}

// /config/network/25-wired.network 修改此文件生效
