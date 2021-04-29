package main

import (
	"example.com/m/v2/cgminer"
	"example.com/m/v2/cgminer/body"
	"log"
	"reflect"
)

//func formatAsDate(t time.Time) string {
//	year, month, day := t.Date()
//	return fmt.Sprintf("%d%02d/%02d", year, month, day)
//}
//
////var netManager *cgminer.NetManager
////var router *gin.Engine
//
//func main22() {
//	netManager = cgminer.CreateInstanc("localhost:4028")
//	router = gin.Default()
//	router.Delims("{[{", "}]}")
//	router.SetFuncMap(template.FuncMap{
//		"formatAsDate": formatAsDate,
//	})
//	router.LoadHTMLFiles("./testdata/template/raw.tmpl")
//	router.GET("/raw", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "raw.tmpl", gin.H{
//			"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
//		})
//	})
//	routerGet()
//	router.Run(":8080")
//}

func main11() {
	pools := make([]body.Pools, 0)
	bodys := make([]cgminer.Body, 0)
	var b cgminer.Body = body.Pools{User: "wdz"}
	var c interface{}
	c = b.(body.Pools)
	bodys = append(bodys, b)
	log.Println(c, reflect.TypeOf(c))
	test(bodys, &pools)
	log.Println(pools)

}

func test(bodys []cgminer.Body, poolsPtr interface{}) {
	log.Println("1", reflect.TypeOf(poolsPtr), reflect.TypeOf(poolsPtr).Elem(), reflect.TypeOf(poolsPtr).Elem().Elem())

	log.Println("2", reflect.TypeOf(bodys), reflect.TypeOf(bodys[0]), reflect.ValueOf(bodys[0]))
	var p interface{}
	p = bodys[0]
	bbb := reflect.ValueOf(p).Convert(reflect.TypeOf(poolsPtr).Elem().Elem())

	log.Println("3", bbb)
	ccc := reflect.Append(reflect.ValueOf(poolsPtr).Elem(), bbb)
	log.Println(ccc)
	reflect.ValueOf(poolsPtr).Elem().Set(ccc)
}
