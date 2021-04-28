package cgminer

import (
	"bytes"
	"encoding/json"
	"example.com/m/v2/cgminer/body"
	"fmt"
	"reflect"
	"strings"
)

type Result struct {
	H Head
	R []*Body
}

//{"STATUS":"S","When":1618385384,"Code":22,"Msg":"CGMiner versions","Description":"cgminer 2.8.7"}
type Head struct {
	STATUS      string
	When        float64
	Code        float64
	Msg         string
	Description string
}
type Body interface {
	Check() bool
	ApiCmd() string
}

var resultMap map[string]interface{}

func Parse(jsonstr string) (*Result, error) {
	var f interface{}
	js := filter0(jsonstr)
	if !json.Valid(*js) {
		fmt.Println("Parse error:json is wrong")
		return nil, JsonError
	}
	err := json.Unmarshal(*js, &f)
	if err != nil {
		fmt.Println("Parse error:", err, " json is=", jsonstr)
		return nil, JsonError
	}
	var head *Head
	var error bool
	var bodys *[]*Body

	m := f.(map[string]interface{})
	for k, v := range m {
		switch v.(type) {
		case []interface{}:
			switch k {
			case "STATUS":
				head, error = getHead(v.([]interface{}))
				if error {
					fmt.Println(k, " head parse failed!!")
					return nil, DataErrorNoHead
				}
			case "id":
				break
			default:
				bodys, error = getBody(k, v.([]interface{}))
				if error {
					fmt.Println(k, " body parse failed!")
					return nil, DataErrorNoBody
				}
			}
		}
	}
	if head == nil {
		fmt.Println("head==nil json is:", jsonstr)
		return nil, DataErrorNoHead
	}
	if bodys == nil || len(*bodys) == 0 {
		fmt.Println("body==nil json is:", jsonstr)
		return nil, DataErrorNoBody
	}
	return &Result{*head, *bodys}, NoError
}

func BodyMap() map[string]interface{} {
	return resultMap
}
func init() {
	resultMap = make(map[string]interface{})
	resultMap["version"] = &body.Version{}
	resultMap["config"] = &body.Config{}
	resultMap["summary"] = &body.Summary{}
	resultMap["pools"] = &body.Pools{}
	resultMap["devs"] = &body.Devs{}
	resultMap["gpu"] = &body.Gpu{}
	resultMap["stats"] = &body.Stats{}
	resultMap["coin"] = &body.Coin{}
	resultMap["check"] = &body.Check{}
	resultMap["devdetails"] = &body.Devdetails{}
}

func filter0(in string) *[]byte {
	a := []byte(in)
	index := bytes.IndexByte(a, 0)
	b := a[:index]
	return &b
}

func getHead(m []interface{}) (*Head, bool) {
	if len(m) != 1 {
		fmt.Println("len(m)!=1 getHead is error")
		return nil, true
	}
	headmap := m[0].(map[string]interface{})
	fmt.Println("head:", headmap)
	_head := reflectNew(Head{})
	for k, v := range headmap {
		value := reflect.ValueOf(v)
		setField(_head, k, value)
	}
	return _head.Interface().(*Head), false
}

func getBody(key string, m []interface{}) (*[]*Body, bool) {
	if len(m) < 1 {
		fmt.Println("len(m)<=1 getBody is error")
		return nil, true
	}
	typeName := strings.ToLower(key)
	tp := resultMap[typeName]
	if tp == nil {
		fmt.Println("getBody:error type=nil,The map can't found the command,key is ", key)
		return nil, true
	}
	bodys := make([]*Body, len(m))
	for index, _ := range m {
		headmap := m[index].(map[string]interface{})
		fmt.Println("body:", headmap)
		_body := reflectNew(tp)
		for k, v := range headmap {
			value := reflect.ValueOf(v)
			setField(_body, k, value)
		}
		fmt.Println("_body:", _body)
		var body Body
		body = _body.Interface().(Body)
		fmt.Println("ptrbody:", body)
		body.Check()
		bodys[index] = &body
	}
	return &bodys, false
}
