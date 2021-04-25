package main

import (
	"encoding/json"
	"example.com/m/v2/cgminer"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
	//sqlite "github.com/mattn/go-sqlite3"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d%02d/%02d", year, month, day)
}

var netManager *cgminer.NetManager
var router *gin.Engine

func main_() {
	netManager = cgminer.CreateInstanc("localhost:4028")
	router = gin.Default()
	router.Delims("{[{", "}]}")
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.LoadHTMLFiles("./testdata/template/raw.tmpl")
	router.GET("/raw", func(c *gin.Context) {
		c.HTML(http.StatusOK, "raw.tmpl", gin.H{
			"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
		})
	})
	routerGet()
	router.Run(":8080")
}

func routerGet() {
	for _, v := range cgminer.BodyMap() {
		body := v.(cgminer.Body)
		router.GET(cgminer.SubUrl+"/"+body.ApiCmd(), func(c *gin.Context) {
			res := make(chan string)
			errorCode := make(chan error)
			fmt.Println("wdz" + body.ApiCmd())
			p := c.DefaultQuery("param", "")
			var data []byte
			if p != "" {
				data, _ = json.Marshal(cgminer.Create(body.ApiCmd(), p))
			} else {
				data, _ = json.Marshal(cgminer.Create(body.ApiCmd(), ""))
			}
			go netManager.TcpCommandSyncByByte(data, res, errorCode)
			r := <-res
			code := <-errorCode
			c.JSON(200, r)
			fmt.Println(code)
			//if code == cgminer.NO_ERROR{
			//	jsonRedult,error:=cgminer.Parse(r)
			//	if error==cgminer.NO_ERROR {
			//		fmt.Println(jsonRedult.H)
			//		for i,v := range jsonRedult.R {
			//			fmt.Println("B",i,"=",*v)
			//		}
			//	}else if code == cgminer.DataErrorNoBody{
			//		fmt.Println(jsonRedult.H)
			//	}
			//}
		})
	}

}
