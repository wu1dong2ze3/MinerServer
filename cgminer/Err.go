package cgminer

import "example.com/m/v2/errs"

var NoError = errs.Definition(0, "No Error")
var JsonError = errs.Definition(-200, "JsonError")
var NetError = errs.Definition(-201, "NetError")
var DataError = errs.Definition(-202, "DataError")
var DataErrorNoHead = errs.Definition(-203, "DataErrorNoHead")
var DataErrorNoBody = errs.Definition(-204, "DataErrorNoBody")
var UnknowError = errs.Definition(-999, "UnknowError")
