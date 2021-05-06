package httpapi

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type MinerHashrate struct {
	BaseJson
	Data MinerHashrateResult `json:"data"`
}
type MinerHashratePost struct {
	SinceTime string `json:"sinceTime"` //"router_type": 1,
	UntilTime string `json:"untilTime"` //"ip": "10.0.0.2",
}

type MinerHashrateResult struct {
	Hashrate HashrateSubResult `json:"hashrate"`
}
type HashrateSubResult struct {
	StartCollectionTime int64    `json:"startCollectionTime"`
	EndCollectionTime   int64    `json:"endCollectionTime"`
	Type                int      `json:"type"`
	Count               int      `json:"count"`
	Points              []Points `json:"points"`
}
type Points struct {
	Index  int       `json:"index"`
	Points []float64 `json:"points"`
}

func (MinerHashrate) GetType() int {
	return POST
}

func (MinerHashrate) GetSubPath() string {
	return "/miner/hashrate"
}

//TODO 假数据
func (MinerHashrate) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		poster := &MinerHashratePost{}
		var err error
		if err = c.ShouldBindJSON(poster); err != nil {
			//c.JSON(http.StatusOK, BaseError(PostJsonError))
			//poster.SinceTime = c.GetInt64("sinceTime")
			//poster.UntilTime = c.GetInt64("untilTime")
		}
		b, _ := json.Marshal(MinerHashrate{*BaseError(NoError),
			MinerHashrateResult{HashrateSubResult{
				time.Now().Unix(),
				time.Now().Unix() + 50000, 0,
				5,
				[]Points{
					{0, []float64{17, 2, 1.1, 0, 5}},
					{1, []float64{12, 1, 2.1, 0, 5}},
					{2, []float64{13, 1, 1.1, 6, 5}},
					{3, []float64{22, 5, 1.3, 0, 5}},
					{-1, []float64{63, 1, 2.1, 0, 111}},
				}}}})
		log.Println(string(b))
		c.JSON(http.StatusOK, MinerHashrate{*BaseError(NoError),
			MinerHashrateResult{HashrateSubResult{
				time.Now().Unix(),
				time.Now().Unix() + 50000, 0,
				5,
				[]Points{
					{0, []float64{17, 2, 1.1, 0, 5}},
					{1, []float64{12, 1, 2.1, 0, 5}},
					{2, []float64{13, 1, 1.1, 6, 5}},
					{3, []float64{22, 5, 1.3, 0, 5}},
					{-1, []float64{63, 1, 2.1, 0, 111}},
				}}}})
	}
}

/**
{
  "code": 200,
  "message": "成功",
  "data": {
    "hashrate": {
      "start_collection_time": 123123432,
      "end_collection_time": 126178431,
      "type": 0,
      "count": 5,
      "points": [
        {
          "index:": 0,
          "points": [
            12.00,
            13.00,
            0,
            0,
            1.0
          ]
        },
        {
          "index:": 1,
          "points": [
            12.00,
            13.00,
            0,
            0,
            1.0
          ]
        },
        {
          "index:": 2,
          "points": [
            12.00,
            13.00,
            0,
            0,
            1.0
          ]
        },
        {
          "index:": 3,
          "points": [
            12.00,
            13.00,
            0,
            0,
            1.0
          ]
        },
        {
          "index:": -1,
          "points": [
            12.00,
            13.00,
            0,
            0,
            1.0
          ]
        }
      ]
    }
  }
}
*/
