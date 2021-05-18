package cgminer

import (
	"example.com/m/v2/cgminer/body"
	"example.com/m/v2/database"
	"example.com/m/v2/errs"
	"log"
	"strconv"
	"sync"
	"time"
)

const DEADLINE = 24 //24小时
const TICKTIME = 5  //5秒采集一次
const SAVERATE = 3  //保存频率 范围1-N 代表过一个采集间隔时间（RATE*TICKTIME）采集一次
type MT struct {
	running   chan bool
	cache     *RCache
	ticker    *time.Ticker
	saveCount int
}

var mt *MT
var onceMT sync.Once

func GetMt() *MT {
	onceMT.Do(func() {
		mt = &MT{}
	})
	return mt
}

func (mt *MT) GetCache() *RCache {
	return mt.cache
}
func (mt *MT) StartLoadMiner() {
	if mt.running != nil {
		return
	}
	mt.ticker = time.NewTicker(TICKTIME * time.Second)
	mt.running = make(chan bool)
	mt.cache = NewRCache()
	log.Println("StartLoadMiner...")
	go func(ticker *time.Ticker) {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				//log.Println("Ticker do")
				mt.doTick()
			case stop := <-mt.running:
				if stop {
					log.Println("Ticker Stop")
					return
				}
			}
		}
	}(mt.ticker)
}

func (mt *MT) Stop() {
	mt.running <- true
	close(mt.running)
	mt.running = nil
}

func (mt *MT) doTick() {
	var errcode *errs.CodeError
	var config *body.Config
	if config, errcode = Config().Get(); errcode != nil {
		log.Println("Error! cgminer config is null! Please check that the CGMiner service is open！ ")
		return
	}
	if config.ASCCount <= 0 {
		log.Println("Error! cgminer ASCCount is Zero! CGMiner service exception！ ")
		return
	}

	var resList *[]body.Asc
	var when int
	if resList, when, errcode = getAscInfo(config.ASCCount); errcode != nil {
		log.Println("Error! getAscInfo error! ")
		return
	}
	//保存一次缓存结果，用于显示
	mt.cache.Sava(resList)
	var points = make([]float64, config.ASCCount)
	for i, v := range *resList {
		points[i] = v.Mhs15M
	}
	//log.Println("online Decice ASCCount:", config.ASCCount)

	if mt.saveCount == 0 {
		//保存数据库结果用于绘图
		//log.Println("database.SavePoints ", time.Now())
		if errcode = database.SavePoints(when, &points); errcode != nil {
			log.Println("SavePoints error!" + errcode.Error())
			return
		}
		//删除过期数据
		if errcode = database.DelPoints((int)(time.Now().Add(-time.Duration(DEADLINE) * time.Hour).Unix())); errcode != nil {
			log.Println("DelPoints error!" + errcode.Error())
			return
		}
		mt.saveCount++
	} else {
		if mt.saveCount+1 >= SAVERATE {
			mt.saveCount = 0
		} else {
			mt.saveCount++
		}
	}

	//测试1分钟过期数据
	//log.Println("time DeadLine:", time.Now().Add(-time.Duration(1)*time.Minute).Unix())
	//if errcode =database.DelPoints((int)(time.Now().Add(-time.Duration(1)*time.Minute).Unix()));errcode!=nil{
	//	log.Println("DelPoints error!" + errcode.Error())
	//	return
	//}
	///测试结束
}
func getAscInfo(ascCount int) (*[]body.Asc, int, *errs.CodeError) {
	var result *Result
	var errCode *errs.CodeError
	var reslist = make([]body.Asc, ascCount)
	var when = 0
	for i := 0; i < ascCount; i++ {
		var oneBody []body.Asc
		if result, errCode = R(body.Asc{}.ApiCmd(), strconv.Itoa(i)); errCode != nil {
			return nil, 0, errCode
		}
		if errCode = result.BindBody(&oneBody); errCode != nil {
			return nil, 0, errCode
		}
		if len(oneBody) != 1 {
			return nil, 0, CGMinerError.AddByString("getAscInfo len(bodyAes)!=1 is " + strconv.Itoa(len(oneBody)))
		}
		reslist[i] = oneBody[0]
		if i == 0 {
			when = result.H.When
		}
	}
	return &reslist, when, nil
}
