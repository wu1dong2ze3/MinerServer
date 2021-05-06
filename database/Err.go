package database

import "example.com/m/v2/errs"

var NoError = errs.Definition(0, "No Error")
var UserPwdError = errs.Definition(-304, "Illegal username and password！")
var UserPwdInvalid = errs.Definition(-305, "Username and password validation failed！")
var PointDatasError = errs.Definition(-306, "Point datas Error！")
var UnknowError = errs.Definition(-999, "UnknowError")
