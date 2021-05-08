package utils

import (
	"example.com/m/v2/errs"
)

var ParamError = errs.Definition(700, "Param Error")
var IpMaskParamError = errs.Definition(701, "Ip or mask ParamError")
