package utils

import (
	"fmt"
	"net"
)

func CheckCidr(ip string, mask string) (string, error) {
	o, all := CheckIP(mask)
	if o == all && all == 0 {
		return "", IpMaskParamError
	}
	var cid = fmt.Sprintf("%s/%d", ip, all)
	_, _, err := net.ParseCIDR(cid)
	if err != nil {
		fmt.Println(err)
		return cid, err
	}
	return "", nil
}

func CheckIP(mask string) (int, int) {
	var m net.IP
	if m = net.ParseIP(mask); m == nil {
		return 0, 0
	}
	ones, bits := net.IPv4Mask(m[0], m[1], m[2], m[3]).Size()
	return ones, bits

}
