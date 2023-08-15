# gormdemo

`GORM` 是 Go 语言中一个强大的 `ORM`（Object-Relational Mapping）库，它提供了方便的方式来操作数据库，将数据库操作转化为面向对象的操作。

该项目作为一个实际示例，演示了如何在 Go 中集成 GORM 以管理数据库操作。

项目目录结构如下：

```shell
gormdemo/
├── main.go
├── models.go
└── utils.go
```

- **gormdemo**：项目的根目录，包含所有的源代码和配置文件。
  - **main.go**：主程序入口文件，包含项目的主函数。这是应用程序的启动点。
  - **models.go**：存放数据模型的定义和数据库操作相关的函数。
  - **utils.go**：包含了通用的工具函数。

# 如何使用

## 1. 克隆项目

使用以下命令在本地计算机上克隆 `gormdemo` 项目：

```shell
git clone https://github.com/zheng-yi-yi/gormdemo.git
```

## 2. 进入项目目录

进入刚刚克隆的项目目录：

```shell
cd gormdemo
```

## 3. 确保已经安装了 Go 以及所需的依赖项

在继续之前，请确保您的系统已经安装了 Go 编程语言，如果还没有安装，您可以从 Go 官方网站 下载并安装 Go。

## 4. 修改 getDSN 函数

在继续之前，请在 utils.go 文件中修改 `getDSN` 函数，以便配置与您的数据库连接相关的信息。将 "username", "password", "your_database_name" 分别替换为实际的用户名、密码和数据库名。

```go
// getDSN 获取数据库连接字符串
func getDSN() string {
	return "username:password@tcp(127.0.0.1:3306)/your_database_name?charset=utf8mb4&parseTime=True&loc=Local"

}
```

## 5. 编译项目并运行

在完成上述步骤后，您可以使用以下命令编译项目并运行应用程序：

```shell
go build
.\gormdemo
```
