package httpapi

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserLogin(t *testing.T) {
	router, _ := InstanceEPM().Test(InstanceRT().GetDefault())
	w := httptest.NewRecorder()
	message := `{"user":"admin","pwd":"admin","device":0}`
	req, _ := http.NewRequest("POST", "/user/login", strings.NewReader(message))
	router.ServeHTTP(w, req)
	respBody, _ := ioutil.ReadAll(w.Body)
	res := UserLogin{}
	err := json.Unmarshal(respBody, &res)
	log.Println("TestUserLogin:", res, err)
	assert.Equal(t, 200, w.Code)
}

var router *gin.Engine = nil

func TestUserUpdate(t *testing.T) {
	TestUserLogin(t)
	router, _ := InstanceEPM().Test(InstanceRT().GetDefault())
	w := httptest.NewRecorder()
	message := `{"user":"admin","pwd":"123456","old_pwd":"admin"}`
	req, _ := http.NewRequest("POST", "/user/update", strings.NewReader(message))
	router.ServeHTTP(w, req)
	respBody, _ := ioutil.ReadAll(w.Body)
	res := UserUpdate{}
	err := json.Unmarshal(respBody, &res)
	log.Println("TestUserUpdate:", res, err)
	assert.Equal(t, 200, w.Code)
}

func TestReboot(t *testing.T) {
	TestUserLogin(t)
	if router == nil {
		router, _ = InstanceEPM().Test(InstanceRT().GetDefault())
	}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/system/reboot", nil)
	req.Header.Add("token", InstanceTKM().Token())
	router.ServeHTTP(w, req)
	respBody, _ := ioutil.ReadAll(w.Body)
	res := BaseJson{}
	err := json.Unmarshal(respBody, &res)
	log.Println("TestReboot:", res, err)
	assert.Equal(t, 200, w.Code)
}

func TestMinerSummary(t *testing.T) {
	TestUserLogin(t)
	if router == nil {
		router, _ = InstanceEPM().Test(InstanceRT().GetDefault())
	}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/miner/summary", nil)
	req.Header.Add("token", InstanceTKM().Token())
	router.ServeHTTP(w, req)
	respBody, _ := ioutil.ReadAll(w.Body)
	res := MinerSummary{}
	err := json.Unmarshal(respBody, &res)
	log.Println("TestMinerSummary:", res, err)
	assert.Equal(t, 200, w.Code)
}

func TestMinerPoolInfo(t *testing.T) {
	TestUserLogin(t)
	if router == nil {
		router, _ = InstanceEPM().Test(InstanceRT().GetDefault())
	}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/miner/pool/info", nil)
	req.Header.Add("token", InstanceTKM().Token())
	router.ServeHTTP(w, req)
	respBody, _ := ioutil.ReadAll(w.Body)
	res := MinerPoolInfo{}
	err := json.Unmarshal(respBody, &res)
	log.Println("TestMinerPoolInfo:", res, err)
	assert.Equal(t, 200, w.Code)
}
func TestMinerHashrate(t *testing.T) {
	TestUserLogin(t)
	if router == nil {
		router, _ = InstanceEPM().Test(InstanceRT().GetDefault())
	}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", MinerHashrate{}.GetSubPath(), nil)
	req.Header.Add("token", InstanceTKM().Token())
	router.ServeHTTP(w, req)
	respBody, _ := ioutil.ReadAll(w.Body)
	res := MinerHashrate{}
	err := json.Unmarshal(respBody, &res)
	log.Println("TestMinerHashrate:", res, err)
	assert.Equal(t, 200, w.Code)
}

func TestMinerBoardInfo(t *testing.T) {
	TestUserLogin(t)
	if router == nil {
		router, _ = InstanceEPM().Test(InstanceRT().GetDefault())
	}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", MinerBoardInfo{}.GetSubPath(), nil)
	req.Header.Add("token", InstanceTKM().Token())
	router.ServeHTTP(w, req)
	respBody, _ := ioutil.ReadAll(w.Body)
	res := MinerBoardInfo{}
	err := json.Unmarshal(respBody, &res)
	log.Println("TestMinerBoardInfo:", res, err)
	assert.Equal(t, 200, w.Code)
}
func TestMinerFansInfo(t *testing.T) {
	TestUserLogin(t)
	if router == nil {
		router, _ = InstanceEPM().Test(InstanceRT().GetDefault())
	}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", MinerFanInfo{}.GetSubPath(), nil)
	req.Header.Add("token", InstanceTKM().Token())
	router.ServeHTTP(w, req)
	respBody, _ := ioutil.ReadAll(w.Body)
	res := MinerFanInfo{}
	err := json.Unmarshal(respBody, &res)
	log.Println("TestMinerFansInfo:", res, err)
	assert.Equal(t, 200, w.Code)
}

func TestMinerUserInfo(t *testing.T) {
	TestUserLogin(t)
	if router == nil {
		router, _ = InstanceEPM().Test(InstanceRT().GetDefault())
	}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", MinerUserInfo{}.GetSubPath(), nil)
	req.Header.Add("token", InstanceTKM().Token())
	router.ServeHTTP(w, req)
	respBody, _ := ioutil.ReadAll(w.Body)
	res := MinerUserInfo{}
	err := json.Unmarshal(respBody, &res)
	log.Println("TestMinerUserInfo:", res, err)
	assert.Equal(t, 200, w.Code)
}

func TestMinerUserUpdate(t *testing.T) {
	TestUserLogin(t)
	if router == nil {
		router, _ = InstanceEPM().Test(InstanceRT().GetDefault())
	}
	info := MinerUserUpdatePost{
		[]UserUpdateSubPost{
			{"address\": \"stratum+tcp://btc.ss.poolin.com:443", "zhiyuan.1", "1"},
			{"address\": \"stratum+tcp://btc.ss.poolin.com:443", "zhiyuan.2", "2"},
			{"address\": \"stratum+tcp://btc.ss.poolin.com:443", "zhiyuan.3", "3"},
		}}
	w := httptest.NewRecorder()
	pbyte, _ := json.Marshal(info)
	req, _ := http.NewRequest("POST", MinerUserUpdate{}.GetSubPath(), strings.NewReader(string(pbyte)))
	req.Header.Add("token", InstanceTKM().Token())
	router.ServeHTTP(w, req)
	respBody, _ := ioutil.ReadAll(w.Body)
	res := MinerUserUpdate{}
	err := json.Unmarshal(respBody, &res)
	log.Println("TestMinerUserUpdate:", res, err)
	assert.Equal(t, 200, w.Code)
}

func TestSystemNetInfo(t *testing.T) {
	TestUserLogin(t)
	if router == nil {
		router, _ = InstanceEPM().Test(InstanceRT().GetDefault())
	}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/system/net/info", nil)
	req.Header.Add("token", InstanceTKM().Token())
	router.ServeHTTP(w, req)
	respBody, _ := ioutil.ReadAll(w.Body)
	res := SystemNetInfo{}
	err := json.Unmarshal(respBody, &res)
	log.Println("TestSystemNetInfo:", res, err)
	assert.Equal(t, 200, w.Code)
}
