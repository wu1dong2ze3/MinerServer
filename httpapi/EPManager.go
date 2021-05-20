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
func (epm EPM) BindEp(eg *gin.Engine) {
	log.Println("Execute")
	eg.MaxMultipartMemory = 1024 * 1024 * 8 //8m
	eg.Use(userAuthorization())
	for _, v := range epArray {
		epm.execute(eg, v)
	}
}

func (epm EPM) Execute(eg *gin.Engine) error {
	log.Println("Execute")
	eg.MaxMultipartMemory = 1024 * 1024 * 8 //8m
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
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token, Accept-Language, Content-Language ,language")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		method := c.Request.Method
		switch method {
		case "OPTIONS":
			c.AbortWithStatus(http.StatusNoContent)
		case "GET", "POST":
			//TODO 正式版本 需要用户权限校验
			//var v = &UserLogin{}
			//if c.Request.URL.Path == v.GetSubPath() {
			//} else {
			//	token := c.GetHeader("token")
			//	log.Println("userAuthorization:", c.Request.URL.Path)
			//	if !InstanceTKM().Check(token) {
			//		log.Println("userAuthorization failed:", c.Request.URL.Path, "token=", token)
			//		c.JSON(http.StatusUnauthorized, *BaseError(ErrToken))
			//		return
			//	}
			//}
		}
		log.Println("userAuthorization", c.FullPath())
		c.Next()
	}
}

var epArray []EndPoint

func init() {
	epArray = []EndPoint{&UserLogin{}, &UserUpdate{}, &SystemReboot{},
		&MinerSummary{}, &MinerPoolInfo{}, &MinerHashrate{}, &MinerBoardInfo{},
		&MinerFanInfo{}, &MinerUserInfo{}, &MinerUserUpdate{}, &SystemNetInfo{},
		&SystemNetUpdate{}, &MinerModeInfo{}, &MinerModeUpdate{}, &SystemOteInfo{},
		&SystemReset{}, &SystemLog{}, &SystemUiTitleBar{}, &SystemHardwareVersion{},
		&SystemHardwareStatus{}, &UserExit{}, &SystemOtaUpgrade{}}
}
