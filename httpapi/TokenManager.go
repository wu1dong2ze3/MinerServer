package httpapi

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"sync"
	"time"
)

type TKM struct {
}
type Token struct {
	User, Pwd string
	time      time.Time
}

const key = "wdztestkey"

var instanceTKM *TKM
var onceTKM sync.Once

var mToken *Token = nil
var mTkString string = ""
var mutex sync.Mutex

func InstanceTKM() *TKM {
	log.Println("InstanceTKM")
	onceTKM.Do(func() {
		instanceTKM = &TKM{}
	})
	return instanceTKM
}

func (tkm TKM) Save(user string, pwd string) (string, error) {
	mutex.Lock()
	defer mutex.Unlock()
	token := &Token{user, pwd, time.Now()}
	mToken = token
	var res string
	var err error
	if res, err = getTokenString(mToken); err != nil {
		return res, err
	}
	mTkString = res
	return res, err

}

func (tkm TKM) Token() string {
	mutex.Lock()
	defer mutex.Unlock()
	return mTkString
}

func (tkm TKM) Check(tkString string) bool {
	mutex.Lock()
	defer mutex.Unlock()
	if mTkString == "" || tkString == "" || mToken == nil {
		log.Println("Check error!", tkString)
		return false
	}
	if tk := parse(tkString); tk != nil {
		if tk.User == mToken.User || tk.Pwd == mToken.Pwd {
			//暂时不校验超时
			return true
		}
	}
	return false
}

func parse(tkString string) *Token {
	token, err := jwt.Parse(tkString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Println("parse", claims["User"], claims["Pwd"], claims["Time"])
		var t time.Time
		if t, err = time.Parse(time.UnixDate, claims["Time"].(string)); err != nil {
			log.Println("Parse token Time error!")
			return nil
		}
		return &Token{claims["User"].(string), claims["Pwd"].(string), t}
	} else {
		fmt.Println(err)
	}
	return nil
}
func getTokenString(token *Token) (string, error) {
	if token == nil {
		return "", NoUser
	}
	var tk = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"User": mToken.User,
		"Pwd":  mToken.Pwd,
		"Time": mToken.time.Format(time.UnixDate),
	})
	tokenString, err := tk.SignedString([]byte(key))
	fmt.Println(tokenString, err)
	return tokenString, err
}
