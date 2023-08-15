package main

import (
	"fmt" // 输入输出标准库

	"gorm.io/driver/mysql" // MySQL数据库驱动
	"gorm.io/gorm"         // 数据库操作
)

func main() {
	// 获取数据库连接字符串
	dsn := getDSN()
	// 连接 MySQL 数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	checkError(err, "连接数据库失败：")
	fmt.Println("数据库连接成功...")

	// 自动创建表结构
	err = db.AutoMigrate(&User{})
	checkError(err, "AutoMigrate()函数运行发生错误：")
	fmt.Println("成功创建表")

	// 创建新用户
	err = createUser(db, "zhangsan", "111111@example.com")
	checkError(err, "无法创建用户1，发生错误：")
	fmt.Println("成功创建用户1...")

	// 创建新用户
	err = createUser(db, "wangwu", "888888@example.com")
	checkError(err, "无法创建用户2，发生错误：")
	fmt.Println("成功创建用户2...")

	// 根据ID查询用户
	user, err := getUserByID(db, 1)
	checkError(err, "无法获取用户1的数据，发生错误：")
	fmt.Println("成功查询用户1，用户名:", user.Username, "邮箱:", user.Email)

	// 更新用户邮箱
	err = updateUserEmail(db, 1, "666666@example.com")
	checkError(err, "无法更新用户1的信息，发生错误：")
	fmt.Println("成功更新用户1的信息...")

	// 删除用户
	err = deleteUser(db, 1)
	checkError(err, "无法删除用户1的信息，发生错误：")
	fmt.Println("成功删除用户1的信息...")

	// 确保在程序结束时关闭数据库连接
	sqlDB, err := db.DB()
	checkError(err, "获取数据库实例失败!")
	defer sqlDB.Close()
}
