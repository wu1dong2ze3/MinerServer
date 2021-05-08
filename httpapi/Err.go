package httpapi

import (
	"example.com/m/v2/errs"
)

var MinerRebootFailed = errs.Definition(670, "Miner reboot failed!")
var NoUser = errs.Definition(671, "User is not logged in") //
var FileUplaodFailed = errs.Definition(672, "FileUplaodFailed")
var CreateWebSocketError = errs.Definition(673, "CreateWebSocketError") //
var WebSocketError = errs.Definition(674, "error! this is a websocket url! ")
var ExeShellError = errs.Definition(675, "Error executing script")         //
var CgMinerDeviceError = errs.Definition(676, "CGMiner devices exception") //
