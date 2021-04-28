package utils

import (
	"log"
	"os"
	"path/filepath"
)

func NowPath() string {
	path, err := os.Executable()
	if err != nil {
		log.Println("NowPath", err)
		return ""
	}
	return filepath.Dir(path)
}
