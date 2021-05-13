package httpapi

import (
	"example.com/m/v2/errs"
	"example.com/m/v2/shell"
	"example.com/m/v2/utils"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
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

const OTA_PATH = "/userdata/ota/"
const UPDATE_FILE_NAME = "update.img"

func (SystemOtaUpgrade) GetHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("SystemOtaUpgrade")
		var err error
		var md5 string
		var baseName string
		var needMd5 = true
		var file *multipart.FileHeader
		//md5 = c.DefaultQuery("md5", "")
		md5, _ = c.GetPostForm("md5")
		if md5 == "" {
			needMd5 = false
		}
		log.Println("md5=", md5)
		if file, err = c.FormFile("file"); err != nil {
			c.JSON(http.StatusOK, *BaseError(errs.IoError.Add(err).AddByString(" FormFile error!Path is " + OTA_PATH)))
			return
		}
		log.Println("fileName:", file.Filename, "size:", file.Size)
		shell.RM.Params("-r " + OTA_PATH).Exec()
		if err = utils.MakeSurePath(OTA_PATH); err != nil {
			c.JSON(http.StatusOK, *BaseError(errs.IoError.Add(err).AddByString(" MakeSurePath error!Path is " + OTA_PATH)))
			return
		}
		baseName = filepath.Base(file.Filename)
		if !strings.Contains(baseName, ".gz") && !strings.Contains(baseName, ".img") && !strings.Contains(baseName, ".tar") {
			c.JSON(http.StatusOK, *BaseError(FileUplaodFailed.AddByString("BaseName is wrong! Only.zip and.img files are supported!" + OTA_PATH + baseName)))
			return
		}
		log.Println("baseName:", baseName)
		if err = c.SaveUploadedFile(file, OTA_PATH+baseName); err != nil {
			c.JSON(http.StatusOK, SystemOtaUpgrade{*BaseError(FileUplaodFailed.Add(err).AddByString("DownlaodFile is wrong!" + OTA_PATH + baseName))})
			return
		}
		if needMd5 {
			if errCode := md5check(OTA_PATH+baseName, md5); errCode != nil {
				c.JSON(http.StatusOK, SystemOtaUpgrade{*BaseError(errCode.AddByString("md5check:"))})
				return
			}
		}
		if strings.Contains(baseName, ".gz") || strings.Contains(baseName, ".tar") {
			//TODO压缩模式
			if _, err = shell.TAR.Params(OTA_PATH+baseName, OTA_PATH).Exec(); err != nil {
				c.JSON(http.StatusOK, SystemOtaUpgrade{*BaseError(errs.IoError.AddByString(" TAR error!file is " + OTA_PATH + baseName))})
				return
			}

		} else {
			if _, err = shell.MV.Params(OTA_PATH + baseName + " " + OTA_PATH + UPDATE_FILE_NAME).Exec(); err != nil {
				c.JSON(http.StatusOK, SystemOtaUpgrade{*BaseError(errs.IoError.AddByString(" MV error!file is " + OTA_PATH + baseName))})
				return
			}
		}
		if errCode := checkFile(OTA_PATH + UPDATE_FILE_NAME); errCode != nil {
			c.JSON(http.StatusOK, SystemOtaUpgrade{*BaseError(errCode.AddByString("checkFile exist:"))})
			return
		}
		//提前返回成功
		c.JSON(http.StatusOK, SystemOtaUpgrade{*BaseError(errs.NoError)})
		go func() {
			//保证服务端收到返回
			time.Sleep(time.Duration(1) * time.Second)
			//shell.STOP_CGMINER.Exec()
			//shell.RECOVERY.Params(OTA_PATH+UPDATE_FILE_NAME).Exec()
		}()

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
