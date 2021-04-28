package httpapi

import (
	"example.com/m/v2/errs"
	"net/http"
)

var NoError = errs.Definition(http.StatusOK, "")
var StatusUnauthorized = errs.Definition(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
var MinerRebootFailed = errs.Definition(427, "Miner reboot failed!")
var PostJsonError = errs.Definition(502, "Request JSON error") //
var NoUser = errs.Definition(503, "User is not logged in")     //
var NoNetConfigFile = errs.Definition(504, "NoNetConfigFile")  //

var CreateWebSocketError = errs.Definition(505, "CreateWebSocketError") //
var WebSocketError = errs.Definition(506, "error! this is a websocket url! ")
var ParamError = errs.Definition(507, "param error! ")
var CanNotLoadFile = errs.Definition(508, "CanNotLoadFile") //
var UnknowError = errs.Definition(999, "UnknowError")
