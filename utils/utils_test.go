package utils

import (
	"example.com/m/v2/utils/simpleticker"
	"fmt"
	"log"
	"testing"
	"time"
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

func TestSimpleTicker(t *testing.T) {

	ss := simpleticker.New()
	ss.DoFunc(func(tickCount int64) {
		fmt.Printf("tcount:%d\n", tickCount)
		if tickCount == 2 {
			ss.Stop()
			ss.Stop()
		}
	}).Start()

	time.Sleep(5 * time.Second)
	fmt.Println("end test1")
	ss.Stop()
	fmt.Println("end test2")

}
func addNumberToChan(chanName chan int) {
	for {
		chanName <- 1
		time.Sleep(1 * time.Second)
	}
}
func TestSimpleTicker2(t *testing.T) {
	var chan1 = make(chan int)
	var chan2 = make(chan int)

	go addNumberToChan(chan1)
	time.Sleep(500 * time.Millisecond)
	go addNumberToChan(chan2)

	for {
		select {
		case e := <-chan1:
			fmt.Printf("Get element from chan1: %d\n", e)
		case e := <-chan2:
			fmt.Printf("Get element from chan2: %d\n", e)
			//default:
			//	fmt.Printf("No element in chan1 and chan2.\n")
			//time.Sleep(0.5 * time.Second)
		}
	}

}
