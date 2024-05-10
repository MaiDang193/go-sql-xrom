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

	//5.engine.ID(1).Update(&user)
	user := User{Name: "maidang"}
	n, _ := engine.ID(1000).Update(&user)
	if n >= 1 {
		fmt.Printf("有%d条数据修改成功！！！", n)
	}

	//6.engine.ID(1).Delete(&user)
	user = User{Name: "maidang"}
	n, _ = engine.ID(1000).Delete(&user)
	if n >= 1 {
		fmt.Printf("有--->%d<---条数据修改成功！！！", n)
	}

	//7.sql语句形式   engine.Exec("update user set age = ? where id = ?", 10, 10001)
	engine.Exec("update user set age = ? where id = ?", 10, 10001)
}
