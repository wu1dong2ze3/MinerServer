package utils

import (
	"encoding/json"
	"example.com/m/v2/errs"
	"fmt"
	"os"
)

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}

func MakeSurePath(path string) error {
	_, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	err = os.Mkdir(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func SaveJson(jsonPtr interface{}, filePathName string) *errs.CodeError {
	var data []byte
	var err error
	if data, err = json.MarshalIndent(jsonPtr, "", "\t"); err != nil {
		return errs.IoError.Add(err)
	}
	if f, err := os.OpenFile(filePathName, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0775); err != nil {
		return errs.IoError.Add(err).AddByString("Create file error！" + filePathName)
	} else {
		defer f.Close()
		if _, err = f.Write(data); err != nil {
			return errs.IoError.Add(err).AddByString("f.Write(data)")
		}
	}

	return nil
}

func LoadJson(jsonPtr interface{}, filePathName string) *errs.CodeError {
	filePtr, err := os.Open(filePathName)
	if err != nil {
		return errs.IoError.AddByString(fmt.Sprintf("Open file failed [Err:%s]\n", err.Error()))
	}
	defer filePtr.Close()
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(jsonPtr)
	if err != nil {
		return errs.IoError.AddByString(fmt.Sprintf("Decoder failed [Err:%s]\n", err.Error()))

	} else {
		return nil
	}
}
