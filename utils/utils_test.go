package utils

import (
	"log"
	"testing"
)

/** dhcp模式
[Match]
Name=eth0
[Network]
DHCP=yes
[DHCP]
ClientIdentifier=mac


[Match]
Name=eth0
[Network]
Netmask=255.255.254.0
Address=192.168.101.231/23
Gateway=192.168.101.240
DNS=192.168.101.250
DNS=233.5.5.5
*/
func TestUserLogin(t *testing.T) {
	ini, _ := New("test.ini")
	ini.Add("[Match]", "")
	ini.Add("Name", "eth0")
	ini.Add("[Network]", "")
	ini.Add("DHCP", "yes")
	ini.Add("[DHCP]", "")
	ini.Add("ClientIdentifier", "mac")
	ini.Save()

	ini, _ = New("test2.ini")
	ini.Add("[Match]", "")
	ini.Add("Name", "eth0")
	ini.Add("[Network]", "")
	ini.Add("Netmask", "255.255.254.0")
	ini.Add("Address", "192.168.101.231/23")
	ini.Add("Gateway", "192.168.101.240")
	ini.Add("DNS", "192.168.101.250")
	ini.Add("DNS", "233.5.5.5")
	ini.Add("DNS", "233.5.5.52")
	ini.Save()

	ini, _ = Open("test2.ini")
	log.Println("TestUserLogin open", ini.LoadByIndex("DNS", 0, "aa"))

}

func TestString(t *testing.T) {
	a, b := CheckCidr("192.168.101.240", "255.255.254.0")

	log.Println("wdz TestString", a, b)

}
