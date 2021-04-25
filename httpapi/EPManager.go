package httpapi

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
)

type EPM struct{}

var instanceEPM *EPM
var onceEPM sync.Once

func InstanceEPM() *EPM {
	log.Println("InstanceEPM")
	onceEPM.Do(func() {
		instanceEPM = &EPM{}
	})
	return instanceEPM
}

func (epm EPM) Execute(eg *gin.Engine) error {
	log.Println("Execute")
	eg.Use(userAuthorization())
	for _, v := range epArray {
		epm.execute(eg, v)
	}
	return eg.Run(":8080")
}
func (epm EPM) Test(eg *gin.Engine) (*gin.Engine, error) {
	log.Println("EPM Test")
	eg.Use(userAuthorization())
	for _, v := range epArray {
		epm.execute(eg, v)
	}
	return eg, nil
}

func (epm EPM) execute(eg *gin.Engine, point EndPoint) *EPM {
	switch point.GetType() {
	case POST:
		eg.POST(point.GetSubPath(), point.GetHandle())
		break
	case GET:
		eg.GET(point.GetSubPath(), point.GetHandle())
		break
	default:
		log.Println("unrealized")
		break
	}
	return &epm
}
func userAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		var v = &UserLogin{}
		if c.Request.URL.Path == v.GetSubPath() {
		} else {
			token := c.GetHeader("token")
			log.Println("userAuthorization:", c.Request.URL.Path)
			if !InstanceTKM().Check(token) {
				log.Println("userAuthorization failed:", c.Request.URL.Path, "token=", token)
				c.JSON(http.StatusUnauthorized, BaseError(StatusUnauthorized))
				return
			}
		}
		c.Next()
	}
}

var epArray []EndPoint

func init() {
	epArray = []EndPoint{&UserLogin{}, &UserUpdate{}, &SystemReboot{},
		&MinerSummary{}, &MinerPoolInfo{}, &MinerHashrate{}, &MinerBoardInfo{},
		&MinerFanInfo{}, &MinerUserInfo{}, &MinerUserUpdate{}, &SystemNetInfo{}}
}
