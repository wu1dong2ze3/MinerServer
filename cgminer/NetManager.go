package cgminer

import (
	"fmt"
	"io/ioutil"
	"net"
)

type NetManager struct {
	addr string
}

var mAddr *NetManager

func (nm NetManager) TcpCommandSyncByByte(data []byte, res chan string, cerr chan error) {
	//获取命令行参数 socket地址
	fmt.Println("TcpCommandSync")
	message := ""
	addr, err := net.ResolveTCPAddr("tcp", mAddr.addr)
	b, _ := checkError(err)
	defer func() {
		res <- message
		cerr <- err
		close(res)
		close(cerr)
	}()
	if b {
		message = ""
		return
	}

	//建立tcp连接
	conn, err := net.DialTCP("tcp", nil, addr)
	b, _ = checkError(err)
	if b {
		message = ""
		return
	}

	//向服务端发送数据
	_, err = conn.Write(data)
	b, _ = checkError(err)
	if b {
		message = ""
		return
	}
	//接收响应
	response, _ := ioutil.ReadAll(conn)
	resRtr := string(response)
	fmt.Println("TcpCommandSyncByByte: ", resRtr)
	message = resRtr
}
func (nm NetManager) TcpCommandSync(cmd string, res chan string, errorCode chan error) {
	nm.TcpCommandSyncByByte([]byte(cmd), res, errorCode)
}

func CreateInstanc(hostAddr string) *NetManager {
	mAddr = &NetManager{hostAddr}
	return mAddr
}

func checkError(err error) (bool, error) {
	errorCode := NoError
	if err == nil {
		return false, errorCode
	}
	switch err {
	default:
		errorCode = NetError
	}
	return true, errorCode
}
