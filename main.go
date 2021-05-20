package main

import (
	"encoding/json"
	"example.com/m/v2/cgminer"
	"example.com/m/v2/database"
	"example.com/m/v2/errs"
	"example.com/m/v2/httpapi"
	"example.com/m/v2/shell"
	"example.com/m/v2/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"strings"
	"time"
)

var nm *cgminer.NM
var router *gin.Engine

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
			c.JSON(200, errs.NoError)
		})
	}

}

type Result struct {
	Address string `json:"address"`
	Port    string `json:"port"`
	Ip      string `json:"ip"`
}

type PortResult struct {
	httpapi.BaseJson
	Data Result `json:"data"`
}

//TODO 代码重构到 factory 里
func htmlRouter() *gin.Engine {
	path := utils.NowPath()
	e := gin.New()
	e.Use(gin.Logger(), gin.Recovery(), func(c *gin.Context) {
		fmt.Println("htmlRouter", c.Request.URL.Path)
		switch c.Request.URL.Path {
		case "/":
			c.Next()
		default:
			if strings.Contains(c.Request.URL.Path, ".") {
				c.Next()
				return
			} else if strings.Contains(c.Request.URL.Path, "/host") {
				c.Header("Access-Control-Allow-Origin", "*")
				c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
				c.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token")
				c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
				c.Header("Access-Control-Allow-Credentials", "true")
				if ip := getIp(); ip == "" {
					fmt.Println("host get ip is null! error!")
					c.JSON(200, *httpapi.BaseError(errs.UnknowError.AddByString("ip is null????")))
				} else {
					log.Println("ip:", ip)
					c.JSON(200, PortResult{httpapi.BaseJson{Code: 200}, Result{ip + ":" + ApiPort, ApiPort, ip}})
				}
				return
			} else {
				fmt.Println("htmlRouter changed " + c.Request.URL.Path + " to " + " /")
				c.Request.URL.Path = "/"
				e.HandleContext(c)
				return
			}
		}
	})
	e.StaticFS("/", http.Dir(path+"/static"))
	return e
}

func apiRouter() *gin.Engine {
	router = httpapi.InstanceRT().GetDefault()
	routerGet()
	httpapi.InstanceEPM().BindEp(router)
	return router
}

var (
	g errgroup.Group
)

const ApiPort = "8080"
const StaticPort = "80"

func main() {

	mt := cgminer.GetMt()
	mt.StartLoadMiner()
	nm = cgminer.GetInstance().Addr("localhost:4028")
	database.GetInstance().DB()
	apiServer := &http.Server{
		Addr:         ":" + ApiPort,
		Handler:      apiRouter(),
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	staticServer := &http.Server{
		Addr:         ":" + StaticPort,
		Handler:      htmlRouter(),
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	g.Go(func() error {
		fmt.Println("GoGohtmlRouter run!")
		return staticServer.ListenAndServe()
	})
	g.Go(func() error {
		fmt.Println("GoGoapiRouter run!")
		return apiServer.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

	mt.Stop()
}
func getIp() string {
	if res, err := shell.ETH0IP.Exec(); err != nil {
		return ""
	} else {
		return utils.S{S: res}.NoSpaceBr().S
	}
}
