package util

import (
	"os"
	"path/filepath"
	"strings"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)

	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}

func GetExeDir() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	res := filepath.Dir(exePath)

	if strings.Contains(exePath, getTmpDir()) {
		// run 模式下，确在当前程序入口目录
		return os.Getwd()
	}
	return res, nil
}

func getTmpDir() string {
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}
	return dir
}
