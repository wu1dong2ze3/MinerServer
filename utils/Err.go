package utils

import (
	"example.com/m/v2/errs"
)

var ParamError = errs.Definition(501, "Param Error")
var IpMaskParamError = errs.Definition(502, "Ip or mask ParamError")
