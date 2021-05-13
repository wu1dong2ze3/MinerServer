package httpapi

import (
	"example.com/m/v2/cgminer"
	"example.com/m/v2/cgminer/body"
	"example.com/m/v2/errs"
	"example.com/m/v2/shell"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

///miner/user/update
type MinerUserUpdate struct {
	BaseJson
}

type MinerUserUpdatePost struct {
	Miners []UserUpdateSubPost `json:"miners"`
}

type UserUpdateSubPost struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Pwd     string `json:"pwd"`
}

func (MinerUserUpdate) GetType() int {
	return POST
}

func (MinerUserUpdate) GetSubPath() string {
	return "/miner/user/update"
}

func (MinerUserUpdate) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		poster := &MinerUserUpdatePost{}
		var err error
		var errCode *errs.CodeError
		// 只支持json模式
		if err = c.ShouldBindJSON(poster); err != nil {
			c.JSON(http.StatusOK, MinerUserUpdate{*BaseError(errs.JsonError)})
			return
		}
		if len(poster.Miners) <= 0 {
			c.JSON(http.StatusOK, MinerUserUpdate{*BaseError(errs.ParamsError.AddByString("len(poster.Miners)<=0"))})
			return
		}
		newPools := make([]cgminer.PoolJson, 0)
		for _, v := range poster.Miners {
			newPools = append(newPools, cgminer.PoolJson{URL: v.Address, User: v.Name, Pass: v.Pwd})
		}
		var poolsInfo = cgminer.PoolsJson{Pools: newPools}
		if errCode = cgminer.InstancePJM().Save(&poolsInfo); errCode != nil {
			c.JSON(http.StatusOK, MinerUserUpdate{*BaseError(errCode)})
			return
		}

		if _, err = shell.RESTART_CGMINER.Exec(); err != nil {
			c.JSON(http.StatusOK, MinerUserUpdate{*BaseError(CgMinerDeviceError.Add(errCode))})
			return
		} else {
			testOk := false
			for i := 0; i < 10; i++ {
				if _, errCode = cgminer.R(body.Version{}.ApiCmd(), ""); errCode != nil {
					time.Sleep(time.Duration(2) * time.Second)
					log.Printf("try to restart! %d times,error is %s", i+1, errCode.Error())
					continue
				}
				testOk = true
				break
			}
			if testOk {
				c.JSON(http.StatusOK, MinerUserUpdate{*BaseError(errs.NoError)})
			} else {
				c.JSON(http.StatusOK, MinerUserUpdate{*BaseError(CgMinerDeviceError.AddByString("cgminer restart over runtime!"))})
			}
			return

		}
		//if _,err=shell.STOP_CGMINER.Exec();err!=nil{
		//	c.JSON(http.StatusOK, MinerUserUpdate{*BaseError(CgMinerDeviceError.Add(err))})
		//	return
		//}
		//if _,err=shell.START_CGMINER.Exec();err!=nil{
		//	c.JSON(http.StatusOK, MinerUserUpdate{*BaseError(CgMinerDeviceError.Add(err))})
		//	return
		//}

	}
}
