package tool

import (
	"go.uber.org/zap"
	"kubernetes_management_system/common"
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

func CreateDirs(dirPaths ...string) (err error) {
	for _, v := range dirPaths {
		exit, err := DirExit(v)
		if err != nil {
			return err
		}
		if !exit {
			common.LOG.Debug("create directory" + v)
			err = os.Mkdir(v, os.ModePerm)
			if err != nil {
				common.LOG.Error("create directory failed"+v, zap.Any("error:", err))
			}
		}
	}
	return err
}

func DeleteDirs(dirPaths ...string) (err error) {
	for _, v := range dirPaths {
		common.LOG.Debug("delete director" + v)
		err := os.Remove(v)
		if err != nil {
			common.LOG.Error("delete directory failed"+v, zap.Any("error:", err))
		}
	}
	return err
}
