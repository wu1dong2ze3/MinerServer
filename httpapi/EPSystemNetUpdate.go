package httpapi

import (
	"example.com/m/v2/errs"
	"example.com/m/v2/shell"
	"example.com/m/v2/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"strconv"
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
		post := &SystemNetUpdatePost{}
		var rType int
		var ip, subnetMask, gateway, dns1, dns2 string
		var err error
		var errCode *errs.CodeError
		if err = c.ShouldBindJSON(post); err != nil {
			rType = c.GetInt("routerType")
			ip = c.GetString("ip")
			subnetMask = c.GetString("subnetMask")
			gateway = c.GetString("gateway")
			dns1 = c.GetString("dns1")
			dns2 = c.GetString("dns2")
		} else {
			rType = post.RouterType
			ip = c.GetString("ip")
			subnetMask = c.GetString("subnetMask")
			gateway = c.GetString("gateway")
			dns1 = c.GetString("dns1")
			dns2 = c.GetString("dns2")
		}
		if errCode = changeNetStatus(rType, ip, subnetMask, gateway, dns1, dns2, func() {
			c.JSON(http.StatusOK, SystemNetUpdate{*BaseError(errs.NoError)})

		}); errCode != nil {
			c.JSON(http.StatusOK, SystemNetUpdate{*BaseError(errCode)})
			return
		}

	}
}

func changeNetStatus(rType int, ip, subnetMask, gateway, dns1, dns2 string, fBeforeClose func()) *errs.CodeError {
	var err error
	var ipcdr string
	//0 static 1 dhcp ,other error
	if rType < 0 || rType > 1 {
		return utils.ParamError.AddByString("rType is " + strconv.Itoa(rType) + ",error!")
	}
	if rType == 0 {
		if ipcdr, err = utils.CheckCidr(ip, subnetMask); err != nil {
			return utils.ParamError.AddByString("ip, subnetMask error!")
		}
		ip = ipcdr
		if gateway != "" {
			if ip := net.ParseIP(gateway); ip == nil {
				return utils.ParamError.AddByString("gateway error! ")
			}
		}
		if dns1 != "" {
			if ip := net.ParseIP(dns1); ip == nil {
				return utils.ParamError.AddByString("Dns1 error!")
			}
		}
		if dns2 != "" {
			if ip := net.ParseIP(dns2); ip == nil {
				return utils.ParamError.AddByString("Dns2 error")
			}
		}
	}
	log.Println("SystemNetUpdatePost", rType, ip, subnetMask, gateway, dns1, dns2)
	//static
	if rType == 0 {
		ini, err := utils.New("/config/network/25-wired.network")
		if err != nil {
			return utils.ParamError.AddByString("New /config/network/25-wired.network error!")
		}
		err = ini.Add("[Match]", "")
		err = ini.Add("Name", "eth0")
		err = ini.Add("[Network]", "")
		err = ini.Add("Netmask", subnetMask)
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
			return utils.ParamError.AddByString("ini.Save() error!")
		}
		if fBeforeClose != nil {
			fBeforeClose()
		}
		_, err = shell.NetworkRestart.Exec()
		if err != nil {
			return utils.ParamError.AddByString("Exec error!")
		}
		//dhcp
	} else if rType == 1 {
		ini, err := utils.New("/config/network/25-wired.network")
		if err != nil {
			return utils.ParamError
		}
		err = ini.Add("[Match]", "")
		err = ini.Add("Name", "eth0")
		err = ini.Add("[Network]", "")
		err = ini.Add("DHCP", "yes")
		err = ini.Add("[DHCP]", "")
		err = ini.Add("ClientIdentifier", "mac")
		if err != nil {
			return utils.ParamError
		}
		if fBeforeClose != nil {
			fBeforeClose()
		}
		_, err = shell.NetworkRestart.Exec()
		if err != nil {
			return utils.ParamError
		}
	}
	return nil
}

// /config/network/25-wired.network 修改此文件生效
