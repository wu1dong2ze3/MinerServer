package main

import (
	"encoding/json"
	"example.com/m/v2/cgminer"
	"example.com/m/v2/httpapi"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

var nm *cgminer.NM
var router *gin.Engine

func main() {
	nm = cgminer.GetInstance().Addr("localhost:4028")
	router = httpapi.InstanceRT().GetDefault()
	routerGet()
	if err := httpapi.InstanceEPM().Execute(router); err != nil {
		log.Println("main", err)
	}
}

func routerGet() {
	for _, v := range cgminer.BodyMap() {
		body := v.(cgminer.Body)
		router.GET(cgminer.SubUrl+"/"+body.ApiCmd(), func(c *gin.Context) {
			fmt.Println("wdz" + body.ApiCmd())
			p := c.DefaultQuery("param", "")
			var data []byte
			if p != "" {
				data, _ = json.Marshal(cgminer.Create(body.ApiCmd(), p))
			} else {
				data, _ = json.Marshal(cgminer.Create(body.ApiCmd(), ""))
			}
			var jsonStr string
			var err error
			if jsonStr, err = nm.TcpCommandSyncByByte(data); err != nil {
				c.JSON(200, err)
				return
			}
			c.JSON(200, jsonStr)
			if jsonR, err := cgminer.Parse(jsonStr); err != nil {
				c.JSON(200, err)
				return
			} else {
				fmt.Println(jsonR)
				for i, v := range jsonR.R {
					fmt.Println("B", i, "=", v)
				}
			}
			c.JSON(200, cgminer.NoError)
		})
	}

}
