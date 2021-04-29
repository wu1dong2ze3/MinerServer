package body

import "strings"

type Devdetails struct {
	Devdetails int    `json:"DEVDETAILS"`
	Name       string `json:"Name"`
	ID         int    `json:"ID"`
	Driver     string `json:"Driver"`
	Kernel     string `json:"Kernel"`
	Model      string `json:"Model"`
	DevicePath string `json:"Device Path"`
}

func (v Devdetails) Check() bool {
	return true
}
func (v Devdetails) ApiCmd() string {
	return strings.ToLower("Devdetails")
}

/**
{
	"STATUS": [{
		"STATUS": "S",
		"When": 1619709545,
		"Code": 69,
		"Msg": "Device Details",
		"Description": "cgminer 4.10.0"
	}],
	"DEVDETAILS": [{
		"DEVDETAILS": 0,
		"Name": "IT1",
		"ID": 0,
		"Driver": "imT1",
		"Kernel": "",
		"Model": "HWmintT1.SingleChain",
		"Device Path": ""
	}, {
		"DEVDETAILS": 1,
		"Name": "IT1",
		"ID": 1,
		"Driver": "imT1",
		"Kernel": "",
		"Model": "HWmintT1.SingleChain",
		"Device Path": ""
	}, {
		"DEVDETAILS": 2,
		"Name": "IT1",
		"ID": 2,
		"Driver": "imT1",
		"Kernel": "",
		"Model": "HWmintT1.SingleChain",
		"Device Path": ""
	}, {
		"DEVDETAILS": 3,
		"Name": "IT1",
		"ID": 3,
		"Driver": "imT1",
		"Kernel": "",
		"Model": "HWmintT1.SingleChain",
		"Device Path": ""
	}],
	"id": 1
}
*/
