package httpapi

import (
	"example.com/m/v2/errs"
	"example.com/m/v2/shell"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

///system/ota/upgrade
type SystemOtaUpgrade struct {
	BaseJson
}

func (SystemOtaUpgrade) GetType() int {
	return POST
}

func (SystemOtaUpgrade) GetSubPath() string {
	return "/system/ota/upgrade"
}

const OTA_PATH = "/userdata/"
const UPDATE_FILE_NAME = "update.img"

func (SystemOtaUpgrade) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("SystemOtaUpgrade")
		var err error
		var md5 string
		var baseName string
		var isZip = false
		var needMd5 = true
		md5 = c.DefaultQuery("md5", "")
		if md5 == "" {
			needMd5 = false
		}
		log.Println("md5=", md5)
		file, _ := c.FormFile("file")
		log.Println("fileName:", file.Filename, "size:", file.Size)
		//TODO 没有写容量判断
		shell.RM.Params(OTA_PATH + UPDATE_FILE_NAME).Exec()
		baseName = filepath.Base(file.Filename)
		if err = c.SaveUploadedFile(file, OTA_PATH+baseName); err != nil {
			c.JSON(http.StatusOK, SystemOtaUpgrade{*BaseError(FileUplaodFailed.Add(err).AddByString("DownlaodFile is wrong!" + OTA_PATH + baseName))})
			return
		}

		if strings.Contains(baseName, ".zip") {
			//TODO压缩模式
			isZip = true
		} else {
			shell.MV.Params(OTA_PATH + baseName + " " + OTA_PATH + UPDATE_FILE_NAME).Exec()
			isZip = false
		}
		if errCode := checkFile(OTA_PATH + UPDATE_FILE_NAME); errCode != nil {
			c.JSON(http.StatusOK, SystemOtaUpgrade{*BaseError(errCode.AddByString("checkFile exist:"))})
			return
		}
		if !isZip && needMd5 {
			if errCode := md5check(OTA_PATH+UPDATE_FILE_NAME, md5); errCode != nil {
				c.JSON(http.StatusOK, SystemOtaUpgrade{*BaseError(errCode.AddByString("md5check:"))})
				return
			}
		}

		//TODO 测试关闭选项
		//shell.STOP_CGMINER.Exec()
		//提前返回
		c.JSON(http.StatusOK, SystemOtaUpgrade{*BaseError(NoError)})
		//shell.RECOVERY.Params(OTA_PATH+UPDATE_FILE_NAME).Exec()
	}
}
func checkFile(file string) *errs.CodeError {
	if _, err := os.Stat(file); err != nil {
		return FileUplaodFailed.Add(err)
	}
	return nil
}

func md5check(file string, md5 string) *errs.CodeError {

	if localMd5, err := shell.MD5SUM.Params(file).Exec(); err != nil {
		return FileUplaodFailed.Add(err)
	} else if !strings.Contains(localMd5, md5) {
		return FileUplaodFailed.AddByString("The wrong md5!" + md5)
	}
	return nil
}
