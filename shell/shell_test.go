package shell

import (
	"log"
	"testing"
)

func TestExec(t *testing.T) {
	//err:=execCommand("find ../ -name \"*.go\"")
	New("find /Users/wudz/bitcoin_device/blockin -name \"*.go\"").ExecCallBack(
		func(out string, err error) (needContinue bool) {
			log.Println(out, err)
			return true
		},
	)

	//assert.Equal(t, true, err, str)
}
