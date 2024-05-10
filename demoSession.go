package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

func main() {
	//1.数据库连接基本信息
	var (
		userName  string = "root"
		password  string = "1234"
		ipAddress string = "127.0.0.1"
		ipPort    int    = 3306
		dbName    string = "go_test"
		charSet   string = "utf8mb4"
	)

	//2.构建数据库连接信息
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddress, ipPort, dbName, charSet)

	fmt.Println("连接名:", dataSourceName)

	//3.创建引擎
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		fmt.Println("数据库连接失败！！！")
		return
	}

	//4.定义一个和表同步的结构体，并且自动同步结构体到数据库
	type User struct {
		Id      int64
		Name    string
		Age     int
		Passwd  string    `xorm:"varchar(200)"`
		Created time.Time `xorm:"created"`
		Updated time.Time `xorm:"updated"`
	}

}
