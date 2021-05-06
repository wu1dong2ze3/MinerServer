package cgminer

import (
	"encoding/json"
	"example.com/m/v2/errs"
	"io/ioutil"
	"log"
	"net"
	"sync"
)

type NM struct {
	addr string
}

var nm *NM
var onceNM sync.Once
var mutex sync.Mutex

func GetInstance() *NM {
	onceNM.Do(func() {
		nm = &NM{}
	})
	return nm
}

func R(cmd string, param string) (*Result, *errs.CodeError) {
	var err error
	var cmdByte []byte
	var resultJson string
	var result *Result
	nm := GetInstance()
	if cmdByte, err = json.Marshal(Create( /*body.Pools{}.ApiCmd()*/ cmd, param)); err != nil {
		return nil, JsonParseError.Add(err).AddByString("json.Marshal error! ")
	}
	if resultJson, err = nm.TcpCommandSyncByByte(cmdByte); err != nil {
		return nil, JsonParseError.Add(err).AddByString("TcpCommandSyncByByte error! ")
	}
	if result, err = Parse(resultJson); err != nil {
		return nil, CGMinerError.Add(err).AddByString("Parse error! ")
	}
	// 测试打印用
	//for i, v := range result.R {
	//	fmt.Println("B", i, "=", v)
	//}
	return result, nil
}

func (nm *NM) tcpCommandByByte(data []byte, res chan string, cerr chan error) {
	message := ""
	var err error = nil
	defer func() {
		res <- message
		cerr <- err
		close(res)
		close(cerr)
	}()
	mutex.Lock()
	defer mutex.Unlock()

	tcpAddr, err := net.ResolveTCPAddr("tcp", nm.addr)
	if err = checkError(err); err != nil {
		err = NoAddrSetError
		return
	}
	//建立tcp连接
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err = checkError(err); err != nil {
		return
	}

	//向服务端发送数据
	_, err = conn.Write(data)
	if err = checkError(err); err != nil {
		return
	}

	//接收响应
	response, _ := ioutil.ReadAll(conn)
	resRtr := string(response)
	//fmt.Println("TcpCommandSyncByByte: ", resRtr)
	message = resRtr
}
func (nm *NM) TcpCommandSync(cmd string) (string, error) {
	return nm.TcpCommandSyncByByte([]byte(cmd))
}
func (nm *NM) TcpCommandSyncByByte(b []byte) (string, error) {
	if nm.addr == "" {
		return "", NoAddrSetError
	}
	result := make(chan string)
	errorCode := make(chan error)
	go nm.tcpCommandByByte(b, result, errorCode)
	r := <-result
	err := <-errorCode
	return r, err
}

func (nm *NM) TcpCommand(cmd string, res chan string, cerr chan error) {
	nm.tcpCommandByByte([]byte(cmd), res, cerr)
}

func (nm *NM) Addr(hostAddr string) *NM {
	nm.addr = hostAddr
	return nm
}

func checkError(err error) error {
	errorCode := NoError
	if err == nil {
		return err
	}
	switch err {
	default:
		errorCode = NetError.Add(err)
	}
	log.Println("checkError;", err)
	return errorCode
}
