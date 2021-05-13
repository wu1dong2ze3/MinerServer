package utils

import (
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
