package utils

import (
	"fmt"
	"os"
)

// 获取命令执行的当前路径
func GetWorkspace() string {
	dir, err := os.Getwd()
	Check(err)
	return dir
}

// 读取配置文件
func ReadConfigFile(w string) []byte {
	filename := fmt.Sprintf("%s%sgdm.xml", w, OsFileSeparator())
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Printf("指定的配置文件不存在：%s", filename)
		return nil
	}
	data, err := os.ReadFile(filename)
	Check(err)
	return data
}

//
