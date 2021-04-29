package cgminer

import (
	"example.com/m/v2/cgminer/body"
	"example.com/m/v2/errs"
	"log"
)

var config *body.Config

func Config() (*body.Config, *errs.CodeError) {
	if config != nil {
		return config, nil
	}
	var errcode *errs.CodeError
	var result *Result
	if result, errcode = R(body.Config{}.ApiCmd(), ""); errcode != nil {
		return nil, errcode
	}
	var bodys = make([]body.Config, 0)
	//需要传指针类型，因为body的内部是这个指针类型的切片
	if errcode = result.BindBody(&bodys); errcode != nil {
		return nil, errcode
	}
	log.Println("bodys=", bodys, len(bodys))
	if len(bodys) != 1 {
		return nil, CGMinerError
	}
	config = &bodys[0]
	return config, nil
}
