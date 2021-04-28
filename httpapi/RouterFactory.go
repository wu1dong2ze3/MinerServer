package httpapi

import (
	"github.com/gin-gonic/gin"
	"sync"
)

type RtFactory struct{}

var singleton *RtFactory
var once sync.Once

func InstanceRT() *RtFactory {
	once.Do(func() {
		singleton = &RtFactory{}
	})
	return singleton
}

func (s RtFactory) GetDefault() *gin.Engine {
	return gin.Default()
}

func (s RtFactory) GetTest() *gin.Engine {
	eg := gin.Default()
	eg.Group("/test")
	return eg
}
