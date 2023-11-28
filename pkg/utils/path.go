package utils

import (
	"errors"
	"io"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				return err
			}
		}
	}
	return err
}

func ReadFile(path string) ([]byte, error) {
	filePtr, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer filePtr.Close()
	return io.ReadAll(filePtr)
}

func WriteFile(path string, data string) error {
	// TODO 此处应该还需要加入目录存在判断
	var d = []byte(data)
	err := os.WriteFile(path, d, os.ModePerm)
	if err != nil {
		return errors.New("write fail")
	}
	return nil
}
