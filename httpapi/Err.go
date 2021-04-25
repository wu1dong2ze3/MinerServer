package httpapi

import (
	"example.com/m/v2/errs"
	"net/http"
)

var NoError = errs.Definition(http.StatusOK, http.StatusText(http.StatusOK))
var StatusUnauthorized = errs.Definition(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
var MinerRebootFailed = errs.Definition(427, "Miner reboot failed!")
var PostJsonError = errs.Definition(502, "Request JSON error") //
var NoUser = errs.Definition(503, "User is not logged in")     //
var UnknowError = errs.Definition(999, "UnknowError")
