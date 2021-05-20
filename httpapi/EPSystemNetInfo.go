package httpapi

import (
	"example.com/m/v2/errs"
	"example.com/m/v2/shell"
	"example.com/m/v2/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
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
	Mac        string `json:"mac"`        //"mac": "38-AD-45-7E-6D-3E",
	RouterType int    `json:"routerType"` //"router_type": 1,
	Ip         string `json:"ip"`         //"ip": "10.0.0.2",
	SubnetMask string `json:"subnetMask"` //"subnet_mask": "255.255.255.0",
	Gateway    string `json:"gateway"`    //"gateway": "10.0.0.1",
	Dns1       string `json:"dns1"`       //"dns1": "114.114.114.114",
	Dns2       string `json:"dns2"`       //"dns2": "202.106.0.20"
}

func (SystemNetInfo) GetType() int {
	return GET
}

func (SystemNetInfo) GetSubPath() string {
	return "/system/net/info"
}

func (SystemNetInfo) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := NetInfoResult{}
		bdns := false

		if ini, error := utils.Open("/config/network/25-wired.network"); error != nil {
			c.JSON(http.StatusOK, *BaseError(errs.IoError.AddByString("/config/network/25-wired.network can't load!")))
			return
		} else {
			if ini.Load("DHCP", "") == "" {
				res.RouterType = 0
			} else {
				res.RouterType = 1
			}
		}
		shell.NetworkStatus.ExecCallBack(func(out string, err error) (needContinue bool) {
			if strings.Contains(out, "HW Address: ") {
				arr := strings.Split(out, ": ")
				log.Println("HW Address:", arr)
				res.Mac = utils.S{S: arr[1]}.NoSpaceBr().S
				return true
			}

			if !strings.Contains(out, "HW Address: ") && strings.Contains(out, "Address: ") {
				arr := strings.Split(out, ": ")
				log.Println("Address:", arr)
				res.Ip = utils.S{S: arr[1]}.NoSpaceBr().S
				return true
			}
			if strings.Contains(out, "Gateway: ") {
				arr := strings.Split(out, ": ")
				log.Println("Gateway:", arr)
				res.Gateway = utils.S{S: arr[1]}.NoSpaceBr().S
				return true
			}
			if strings.Contains(out, "DNS: ") {
				arr := strings.Split(out, ": ")
				log.Println("DNS1:", arr)
				res.Dns1 = utils.S{S: arr[1]}.NoSpaceBr().S
				bdns = true
				return true
			}
			if bdns == true {
				log.Println("DNS2:", out)
				res.Dns2 = utils.S{S: out}.NoSpaceBr().S
				bdns = false
				return true
			}
			return true
		})
		beth0 := false
		shell.IFCONFIG.ExecCallBack(func(out string, err error) (needContinue bool) {
			if strings.Contains(out, "eth0") {
				beth0 = true
				return true
			}
			if beth0 {
				index := strings.LastIndex(out, "Mask:")
				if index != -1 {
					ar := strings.Split(out, ":")
					log.Println("splite:", ar, len(ar))
					if len(ar) >= 4 {
						res.SubnetMask = utils.S{S: ar[3]}.NoSpaceBr().S
					}
				}
				beth0 = false
				return false
			}
			return true
		})

		log.Println("/system/net/info", res)
		c.JSON(http.StatusOK, SystemNetInfo{BaseJson{http.StatusOK, ""},
			SystemNetInfoResult{
				res}})
	}
}

// /config/network/25-wired.network 修改此文件生效

/** dhcp模式
[Match]
Name=eth0
[Network]
DHCP=yes
[DHCP]
ClientIdentifier=mac
*/
/*
[Match]
Name=eth0
[Network]
Netmask=255.255.254.0
Address=192.168.101.231/23
Gateway=192.168.101.240
DNS=192.168.101.250
DNS=233.5.5.5
*/
