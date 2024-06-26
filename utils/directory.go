package utils

import (
	"github.com/gaochuang/cloudManagementSystem/pkg/log"
	"go.uber.org/zap"
	"os"
)

func DirExit(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func IsDirectoryExists(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

func CreateDirs(dirPaths ...string) (err error) {
	for _, v := range dirPaths {
		exit, err := DirExit(v)
		if err != nil {
			return err
		}
		if !exit {
			err = os.Mkdir(v, os.ModePerm)
			if err != nil {
				log.Logger.LogError("create directory failed"+v, zap.Any("error:", err))
			}
		}
	}
	return err
}

func DeleteDirs(dirPaths ...string) (err error) {
	for _, v := range dirPaths {
		log.Logger.LogInfo("delete director" + v)
		err := os.Remove(v)
		if err != nil {
			log.Logger.LogError("delete directory failed"+v, zap.Any("error:", err))
		}
	}
	return err
}

func EnsureDirectoryExists(path string) error {
	if !IsDirectoryExists(path) {
		if err := CreateDirs(path); err != nil {
			return err
		}
	}
	return nil
}
