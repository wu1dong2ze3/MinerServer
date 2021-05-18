package utils

import (
	"example.com/m/v2/errs"
	"example.com/m/v2/shell"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func NowPath() string {
	path, err := os.Executable()
	if err != nil {
		log.Println("NowPath", err)
		return ""
	}
	return filepath.Dir(path)
}

//获取固件相关信息
const configFile = "/etc/sw_info"

func GetVersionInfo() (string, string, *errs.CodeError) {
	if !IsExist(configFile) {
		return "", "", errs.IoError.AddByString(configFile + " is not exist!")
	}
	var err error
	var ini *Ini
	if ini, err = Open(configFile); err != nil {
		return "", "", errs.IoError.AddByString(configFile + " can't open!")
	}
	var version string
	if version = ini.Load("sw_version", ""); version == "" {
		return "", "", errs.IoError.AddByString("sw_version is null")
	}

	var timeStr string
	shell.STAT.Params(configFile).ExecCallBack(func(out string, err error) (needContinue bool) {
		if strings.Contains(out, "Modify:") {
			timeStr = out
			timeStr = S{timeStr}.NoAny("Modify:").NoBr().Select(" ", " ", 1, 2).Select(".", "", 0).S
			return false
		}
		return true
	})
	return version, timeStr, nil
}
