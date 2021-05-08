package database

import "example.com/m/v2/errs"

var UserPwdError = errs.Definition(650, "Illegal username and password！")
var PointDatasError = errs.Definition(651, "Point datas Error！")
