package utils

import (
	"fmt"
	"log"
	"net"
)

func CheckCidr(ip string, mask string) (string, error) {
	o, all := CheckIP(mask)
	log.Println("wdz CheckCidr", o, all, " ip="+ip)
	if o == all && all == 0 {
		return "", IpMaskParamError
	}
	var cid = fmt.Sprintf("%s/%d", ip, o)
	_, _, err := net.ParseCIDR(cid)
	if err != nil {
		fmt.Println("CheckCidr error:", err)
		return cid, err
	}
	return cid, nil
}

func CheckIP(mask string) (int, int) {
	var m net.IP
	if m = net.ParseIP(mask); m == nil {
		return 0, 0
	}
	if len(m) != 16 {
		return 0, 0
	}
	ones, bits := net.IPv4Mask(m[12], m[13], m[14], m[15]).Size()
	//	ones = 0
	//
	//ALL:
	//	for i := 0; i < 4; i++ {
	//		var s byte = 0x80
	//		for j := 0; j < 8; j++ {
	//			if s&m[i] != 0 {
	//				log.Println("CheckIP=", s, m[i])
	//				ones++
	//				s=s<<1
	//			} else {
	//				break ALL
	//			}
	//		}
	//	}
	log.Println("CheckIP=", ones, bits)
	if ones == 0 {
		return 0, 0
	}
	return ones, bits

}
