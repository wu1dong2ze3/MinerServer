package cgminer

import "example.com/m/v2/errs"

var NetError = errs.Definition(601, "NetError")
var DataError = errs.Definition(602, "DataError")
var DataErrorNoHead = errs.Definition(603, "DataErrorNoHead")
var DataErrorNoBody = errs.Definition(604, "DataErrorNoBody")
var NoAddrSetError = errs.Definition(605, "No address set")
var BodyDataError = errs.Definition(606, "BodyDataError")
var CGMinerError = errs.Definition(607, "CGMinerError!")
