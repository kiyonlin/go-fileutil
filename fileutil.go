package fileutil

import (
	"os"
	"path/filepath"
)

// CreateFile 创建文件
func CreateFile(filePath string) (*os.File, error) {
	if path, err := MakeFilePath(filepath.Dir(filePath), filepath.Base(filePath)); err != nil {
		return nil, err
	} else {
		return os.Create(path)
	}
}

// MakeFilePath 文件路径，当目录不存在时，进行创建
func MakeFilePath(dirName, fileName string) (string, error) {
	if err := EnsureDir(dirName); err != nil {
		return "", err
	}
	return filepath.Join(dirName, fileName), nil
}

// EnsureDir 保证目录存在
func EnsureDir(dirName string, mode ...os.FileMode) error {
	m := os.FileMode(0750)
	if len(mode) > 0 {
		m = mode[0]
	}

	if err := os.MkdirAll(dirName, m); err == nil || os.IsExist(err) {
		return nil
	} else {
		return err
	}
}

// FileExists 判断文件存不存在
func FileExists(path string) bool {
	// os.Stat获取文件信息
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
