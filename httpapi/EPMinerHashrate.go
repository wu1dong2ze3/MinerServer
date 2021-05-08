package httpapi

import (
	"example.com/m/v2/cgminer"
	"example.com/m/v2/database"
	"example.com/m/v2/errs"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type MinerHashrate struct {
	BaseJson
	Data MinerHashrateResult `json:"data"`
}
type MinerHashratePost struct {
	SinceTime int `json:"sinceTime"` //"router_type": 1,
	UntilTime int `json:"untilTime"` //"ip": "10.0.0.2",
}

type MinerHashrateResult struct {
	Hashrate HashrateSubResult `json:"hashrate"`
}
type HashrateSubResult struct {
	SinceTime int      `json:"sinceTime"`
	UntilTime int      `json:"untilTime"`
	Type      int      `json:"type"`
	ResCount  int      `json:"resCount"`
	Points    []Points `json:"points"`
}
type Points struct {
	Index  int     `json:"index"`
	Count  int     `json:"count"`
	Points []Point `json:"points"`
}
type Point struct {
	X int     `json:"x"`
	Y float64 `json:"y"`
}

func (MinerHashrate) GetType() int {
	return POST
}

func (MinerHashrate) GetSubPath() string {
	return "/miner/hashrate"
}

func (MinerHashrate) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		poster := &MinerHashratePost{}
		var err error
		var codeErr *errs.CodeError
		if err = c.ShouldBindJSON(&poster); err != nil {
			//保留 type=0
			poster.SinceTime = c.GetInt("sinceTime")
			poster.UntilTime = c.GetInt("untilTime")
		}
		var deviceCount = 0
		if config, errCode := cgminer.Config().Get(); errCode != nil {
			c.JSON(http.StatusOK, *BaseError(errCode))
			return
		} else {
			deviceCount = config.ASCCount
			if deviceCount == 0 {
				//如果设备为空返回空结构体json
				c.JSON(http.StatusOK, MinerHashrate{})
				return
			}
		}

		var pointList *[]database.PointList
		var pointsCount int
		if pointList, codeErr = database.QueryPoints(poster.SinceTime, poster.UntilTime); codeErr != nil {
			c.JSON(http.StatusOK, *BaseError(codeErr))
			return
		}
		pointsCount = len(*pointList)
		if pointsCount == 0 {
			//如果数据库空返回空结构体json
			c.JSON(http.StatusOK, MinerHashrate{})
			return
		}
		//数据库的len必须大于实际ASCCount 否则数组溢出
		if deviceCount > database.PointsCapacity {
			c.JSON(http.StatusOK, *BaseError(CgMinerDeviceError.AddByString("MinerHashrate:len(*pointList)<count or count > database.PointsCapacity")))
			return
		}
		mhr := MinerHashrate{*BaseError(errs.NoError), MinerHashrateResult{HashrateSubResult{}}}
		mhr.Data.Hashrate.Points = make([]Points, deviceCount+1)
		mhr.Data.Hashrate.SinceTime, mhr.Data.Hashrate.UntilTime = getTime(pointList)
		mhr.Data.Hashrate.ResCount = deviceCount + 1 //需要加上总结果数量
		////拼装普通采集点
		for i := 0; i < deviceCount; i++ {
			getPoints(i, &mhr.Data.Hashrate.Points[i], pointList)
		}
		////获取全部总算力采集点
		getTotal(&mhr.Data.Hashrate.Points[deviceCount], pointList)
		c.JSON(http.StatusOK, mhr)
	}
}

func getTime(dbList *[]database.PointList) (sinceTime int, untilTime int) {
	if len(*dbList) > 0 {
		return (*dbList)[0].Time, (*dbList)[len(*dbList)-1].Time
	} else {
		return 0, 0
	}
}
func getPoints(index int, ptrPoints *Points, dbList *[]database.PointList) {
	ptrPoints.Index = index
	ptrPoints.Count = len(*dbList)
	ptrPoints.Points = make([]Point, ptrPoints.Count)
	for i, v := range *dbList {
		ptrPoints.Points[i].Y = reflectFloat64(index+1, v)
		ptrPoints.Points[i].X = v.Time
	}
}
func reflectFloat64(index int, point database.PointList) float64 {
	fieldName := fmt.Sprintf("Point%d", index)
	return reflect.ValueOf(point).FieldByName(fieldName).Float()
}
func getTotal(ptrPoints *Points, dbList *[]database.PointList) {
	ptrPoints.Index = -1
	ptrPoints.Count = len(*dbList)
	ptrPoints.Points = make([]Point, ptrPoints.Count)
	for i, v := range *dbList {
		ptrPoints.Points[i].Y = v.PointTotal
		ptrPoints.Points[i].X = v.Time
	}
}

/**
{
	"code": 200,
	"message": "",
	"data": {
		"hashrate": {
			"sinceTime": 1620466053,
			"untilTime": 1620466063,
			"type": 0,
			"resCount": 5,
			"points": [{
				"index": 0,
				"count": 3,
				"points": [{
					"x": 1620466053,
					"y": 42850.28
				}, {
					"x": 1620466058,
					"y": 42850.34
				}, {
					"x": 1620466063,
					"y": 42850.38
				}]
			}, {
				"index": 1,
				"count": 3,
				"points": [{
					"x": 1620466053,
					"y": 42850.26
				}, {
					"x": 1620466058,
					"y": 42850.31
				}, {
					"x": 1620466063,
					"y": 42850.36
				}]
			}, {
				"index": 2,
				"count": 3,
				"points": [{
					"x": 1620466053,
					"y": 42850.26
				}, {
					"x": 1620466058,
					"y": 42850.32
				}, {
					"x": 1620466063,
					"y": 42850.37
				}]
			}, {
				"index": 3,
				"count": 3,
				"points": [{
					"x": 1620466053,
					"y": 42850.26
				}, {
					"x": 1620466058,
					"y": 42850.31
				}, {
					"x": 1620466063,
					"y": 42850.36
				}]
			}, {
				"index": -1,
				"count": 3,
				"points": [{
					"x": 1620466053,
					"y": 171401.06000000003
				}, {
					"x": 1620466058,
					"y": 171401.28
				}, {
					"x": 1620466063,
					"y": 171401.46999999997
				}]
			}]
		}
	}
}
*/
