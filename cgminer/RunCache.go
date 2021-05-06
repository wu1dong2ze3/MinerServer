package cgminer

import (
	"example.com/m/v2/cgminer/body"
	"example.com/m/v2/errs"
	"log"
	"sync"
	"time"
)

/**
单例：
缓存一些运行中基础信息，省去拿一些不必要的网络数据
*/
type ConfigCache struct {
	lock sync.Mutex
}

var cCache *ConfigCache
var onceRCache sync.Once

func Config() *ConfigCache {
	onceRCache.Do(func() {
		cCache = &ConfigCache{}
	})
	return cCache
}

/**
运行时期缓存的 Config数据
*/
var config *body.Config

func (r *ConfigCache) Get() (*body.Config, *errs.CodeError) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if config != nil {
		return config, nil
	}
	var errcode *errs.CodeError
	var result *Result
	if result, errcode = R(body.Config{}.ApiCmd(), ""); errcode != nil {
		return nil, errcode
	}
	var bodys = make([]body.Config, 0)
	//需要传指针类型，因为body的内部是这个指针类型的切片
	if errcode = result.BindBody(&bodys); errcode != nil {
		return nil, errcode
	}
	log.Println("bodys=", bodys, len(bodys))
	if len(bodys) != 1 {
		return nil, CGMinerError
	}
	config = &bodys[0]
	return config, nil
}

/**
运行时期缓存的各种数据，以时间为条件去判断是否可用
*/
type RCache struct {
	cache  interface{}
	lsTime time.Time
	lock   sync.Mutex
}

func NewRCache() *RCache {
	return &RCache{nil, time.Now(), sync.Mutex{}}
}

func (r *RCache) Sava(c interface{}) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.cache = c
	r.lsTime = time.Now()
	//log.Println("SavaTemp:",c,r.lsTime)
}

func (r *RCache) Load(d time.Duration) (interface{}, *errs.CodeError) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.cache == nil {
		return nil, errs.New("Cache is nil", 1001)
	}

	if r.lsTime.Add(d).After(time.Now()) {
		return r.cache, nil
	}
	return r.cache, errs.New("Over run Time!", 1002)
}
