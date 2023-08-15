package main

import (
	"fmt"
	"os"
)

// getDSN 获取数据库连接字符串
func getDSN() string {
	return "username:password@tcp(127.0.0.1:3306)/your_database_name?charset=utf8mb4&parseTime=True&loc=Local"

}

// checkError 检查错误并进行处理
func checkError(err error, message string) {
	if err != nil {
		fmt.Println(message, err)
		os.Exit(1)
	}
}
