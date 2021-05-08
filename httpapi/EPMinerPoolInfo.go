package httpapi

import (
	"example.com/m/v2/cgminer"
	"example.com/m/v2/cgminer/body"
	"example.com/m/v2/errs"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type MinerPoolInfo struct {
	BaseJson
	Data MinerPoolInfoResult `json:"data"`
}

type MinerPoolInfoResult struct {
	Pools []PoolInfoSubResult `json:"pools"`
}

type PoolInfoSubResult struct {
	Index      int    `json:"index"`      //"index": 0,
	Address    string `json:"address"`    //"address": "stratum+tcp://btc.ss.poolin.com:443",
	Status     int    `json:"status"`     //"status": 1,
	Name       string `json:"name"`       //"name": "zhiyuan.5x36",
	Diff       string `json:"diff"`       //"diff": "262K",
	ReceiveNum int    `json:"receiveNum"` //"receive_num": 2344,
	RejectNum  int    `json:"rejectNum"`  //"reject_num": 1,
	LsTime     int64  `json:"lsTime"`     //"ls_time": 1663445224
}

func (MinerPoolInfo) GetType() int {
	return GET
}

func (MinerPoolInfo) GetSubPath() string {
	return "/miner/pool/info"
}

func (MinerPoolInfo) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var errcode *errs.CodeError
		var result *cgminer.Result
		if result, errcode = cgminer.R(body.Pools{}.ApiCmd(), ""); errcode != nil {
			c.JSON(http.StatusOK, *BaseError(errcode))
			return
		}
		var bodys = make([]body.Pools, 0)
		//需要传指针类型，因为body的内部是这个指针类型的切片
		if errcode = result.BindBody(&bodys); errcode != nil {
			c.JSON(http.StatusOK, *BaseError(cgminer.CGMinerError.Add(errcode)))
			return
		}
		log.Println("bodys=", bodys, len(bodys))
		resps := make([]PoolInfoSubResult, 0)
		for _, v := range bodys {
			resps = append(resps, PoolInfoSubResult{
				Index:      v.Pool,
				Address:    v.URL,
				Status:     cgminer.StatusType(v.Status),
				Name:       v.User,
				Diff:       fmt.Sprintf("%.2fK", v.WorkDifficulty),
				ReceiveNum: v.Accepted,
				RejectNum:  v.Rejected,
				LsTime:     int64(v.LastShareTime),
			})
		}

		c.JSON(http.StatusOK, MinerPoolInfo{*BaseError(errs.NoError),
			MinerPoolInfoResult{resps}})
	}
}

/**
{
	"code": 200,
	"message": "",
	"data": {
		"pools": [{
			"index": 1,
			"address": "stratum+tcp://btc.ss.poolin.com:443",
			"status": 0,
			"name": "vansbtc.001",
			"diff": "8192.00K",
			"receiveNum": 1,
			"rejectNum": 2,
			"lsTime": 0
		}, {
			"index": 2,
			"address": "stratum+tcp://btc.ss.poolin.com:25",
			"status": 0,
			"name": "vansbtc.002",
			"diff": "65536.00K",
			"receiveNum": 0,
			"rejectNum": 0,
			"lsTime": 0
		}, {
			"index": 3,
			"address": "stratum+tcp://btc.ss.poolin.com:1883",
			"status": 0,
			"name": "vansbtc.003",
			"diff": "65536.00K",
			"receiveNum": 0,
			"rejectNum": 0,
			"lsTime": 0
		}]
	}
}
*/
