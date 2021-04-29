package cgminer

import "example.com/m/v2/errs"

var NoError = errs.Definition(0, "No Error")
var JsonError = errs.Definition(-600, "JsonError")
var NetError = errs.Definition(-601, "NetError")
var DataError = errs.Definition(-602, "DataError")
var DataErrorNoHead = errs.Definition(-603, "DataErrorNoHead")
var DataErrorNoBody = errs.Definition(-604, "DataErrorNoBody")
var NoAddrSetError = errs.Definition(-606, "No address set")
var BodyDataError = errs.Definition(-607, "BodyDataError")
var UnknowError = errs.Definition(-999, "UnknowError")
var JsonParseError = errs.Definition(608, "Parse json error!")
var CGMinerError = errs.Definition(609, "CGMinerError!")
