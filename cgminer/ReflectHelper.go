package cgminer

import (
	"fmt"
	"reflect"
	"strings"
)

func reflectNew(target interface{}) *reflect.Value {
	if target == nil {
		fmt.Println("参数不能未空")
		return nil
	}
	t := reflect.TypeOf(target)
	if t.Kind() == reflect.Ptr { //指针类型获取真正type需要调用Elem
		t = t.Elem()
		fmt.Println("reflectNew t.Kind() == reflect.Ptr")
	}
	newStruc := reflect.New(t)
	return &newStruc
}

func setField(newStruc *reflect.Value, key string, v reflect.Value) {
	//fmt.Println("setField:", key, " v:", v, " v.type:", v.Type())
	tpvalue := newStruc.Elem().Type()
	//fmt.Println("setField:ntpvalue", tpvalue)
	b := false
	for i := 0; i < tpvalue.NumField(); i++ {
		if strings.Contains(string(tpvalue.Field(i).Name), key) || strings.Contains(string(tpvalue.Field(i).Tag), key) {
			vv := newStruc.Elem().Field(i)
			//float64 对 json int 的问题
			if "int" == tpvalue.Field(i).Type.Name() {
				iv := int(v.Float())
				vv.Set(reflect.ValueOf(iv))
			} else {
				vv.Set(v)
			}
			b = true
			break
		}
	}
	if !b {
		fmt.Println("setField:warning !!! the key=", key, " is not found!")
	}

}
