package main

import (
	"gorm.io/gorm"
)

// User 是一个数据模型结构
type User struct {
	ID       uint   `gorm:"primaryKey"`      // 定义用户的唯一标识，并使用 gorm 的 primaryKey 标签指定为主键
	Username string `gorm:"column:username"` // 定义用户名字段，并使用 gorm 的 column 标签指定数据库列名为 "username"
	Email    string `gorm:"column:email"`    // 定义电子邮件地址字段，并使用 gorm 的 column 标签指定数据库列名为 "email"
}

// createUser 创建一个新用户记录
func createUser(db *gorm.DB, username, email string) error {
	// 创建一个新的 User 结构体实例，并初始化其中的字段
	newUser := &User{
		Username: username, // 将传入的用户名赋值给 newUser 结构体的 Username 字段
		Email:    email,    // 将传入的电子邮件地址赋值给 newUser 结构体的 Email 字段
	}

	// 使用 GORM 的 Create 方法插入记录
	result := db.Create(newUser) // 在数据库中插入新用户记录
	return result.Error          // 返回可能出现的错误
}

// getUserByID 根据用户 ID 查询用户记录
func getUserByID(db *gorm.DB, id uint) (*User, error) {
	var user User
	result := db.First(&user, id) // 在数据库中查找指定 ID 的用户记录
	if result.Error != nil {
		return nil, result.Error // 如果出现错误，返回 nil 和错误信息
	}
	return &user, nil // 返回找到的用户记录和 nil 错误
}

// updateUserEmail 更新用户的电子邮件地址
func updateUserEmail(db *gorm.DB, id uint, newEmail string) error {
	user := &User{}
	result := db.First(user, id) // 在数据库中查找指定 ID 的用户记录
	if result.Error != nil {
		return result.Error // 如果出现错误，返回错误信息
	}

	// 使用 GORM 的 Model 和 Update 方法更新记录字段
	result = db.Model(user).Update("Email", newEmail) // 更新用户记录中的电子邮件字段
	return result.Error                               // 返回可能出现的错误
}

// deleteUser 根据用户 ID 删除用户记录
func deleteUser(db *gorm.DB, id uint) error {
	var user User
	result := db.First(&user, id) // 在数据库中查找指定 ID 的用户记录
	if result.Error != nil {
		return result.Error // 如果出现错误，返回错误信息
	}

	// 使用 GORM 的 Delete 方法删除记录
	result = db.Delete(&user) // 从数据库中删除用户记录
	return result.Error       // 返回可能出现的错误
}
