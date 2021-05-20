package simpleticker

import (
	"errors"
	"log"
	"time"
)

/**
非阻塞定时器
*/

var (
	ErrInitialization = errors.New("incorrect initialization parameter")
	ErrTickerRunning  = errors.New("ticker is running")
)

type STicker struct {
	running  chan bool
	ticker   *time.Ticker
	interval time.Duration
	funcPtr  func(tickCount int64)
}

func New() *STicker {
	return &STicker{interval: 1 * time.Second}
}

func (st *STicker) Interval(d time.Duration) *STicker {
	st.interval = d
	return st
}

func (st *STicker) DoFunc(funcPtr func(tickCount int64)) *STicker {
	st.funcPtr = funcPtr
	return st
}

func (st *STicker) Start() error {
	if st.funcPtr == nil {
		return ErrInitialization
	}
	if st.ticker != nil {
		return ErrTickerRunning
	}
	st.ticker = time.NewTicker(st.interval)
	//防止Dofunc内部调用stop阻塞线程 这里申请2个信号量
	st.running = make(chan bool, 2)
	log.Println("Start...")
	var count int64
	go func(ticker *time.Ticker) {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				st.funcPtr(count)
				count++
			case stop := <-st.running:
				log.Println("Ticker Stop", stop)
				return
			}
		}
	}(st.ticker)
	return nil
}

func (st *STicker) Stop() {
	if st.ticker != nil {
		st.running <- true
		close(st.running)
		st.ticker = nil
	}
}
