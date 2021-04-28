package shell

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

type CallBack func(out string, err error) (needContinue bool)

type Cmd struct {
	cmd string
}

func New(cmd string) *Cmd {
	return &Cmd{cmd}
}

//阻塞式
func (c Cmd) Exec() (string, error) {
	cmd := exec.Command("/bin/bash", "-c", c.cmd)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

//阻塞式通过回调处理返回
func (c Cmd) ExecCallBack(callBack CallBack) {
	cmd := exec.Command("/bin/bash", "-c", c.cmd)
	fmt.Println(cmd.Args)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		callBack("", err)
		return
	}
	if err = cmd.Start(); err != nil {
		callBack("", err)
		return
	}
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if !callBack(line, err2) {
			break
		}
		if io.EOF == err2 {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		callBack("", err)
		return
	}
	fmt.Println("execCommand", "end")
}

func (c Cmd) Params(params ...string) *Cmd {
	for _, v := range params {
		c.cmd = c.cmd + " " + v
	}
	return &c
}

//阻塞式通过回调处理返回
func (c Cmd) ExecCallBackNoEOF(callBack CallBack) {
	cmd := exec.Command("/bin/bash", "-c", c.cmd)
	fmt.Println(cmd.Args)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		callBack("", err)
		return
	}
	if err = cmd.Start(); err != nil {
		callBack("", err)
		return
	}
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if !callBack(line, err2) {
			break
		}
		//除非主动退出，否则一直阻断
		//if io.EOF == err2 {
		//	break
		//}
	}
	if err = cmd.Wait(); err != nil {
		callBack("", err)
		return
	}
	fmt.Println("execCommand", "end")
}

//不需要执行命令的结果与成功与否，执行命令马上就返回
func exec_shell_no_result(command string) {
	////处理启动参数，通过空格分离 如：setsid /home/luojing/gotest/src/test_main/iwatch/test/while_little &
	//command_name_and_args := strings.FieldsFunc(command, splite_command)
	////开始执行c包含的命令，但并不会等待该命令完成即返回
	//cmd.Start()
	//if err != nil {
	//	fmt.Printf("%v: exec command:%v error:%v\n", get_time(), command, err)
	//}
	//fmt.Printf("Waiting for command:%v to finish...\n", command)
	////阻塞等待fork出的子进程执行的结果，和cmd.Start()配合使用[不等待回收资源，会导致fork出执行shell命令的子进程变为僵尸进程]
	//err = cmd.Wait()
	//if err != nil {
	//	fmt.Printf("%v: Command finished with error: %v\n", get_time(), err)
	//}
	//return
}
