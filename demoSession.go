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

	//5.事务应用
	//5.1.构建事务对象
	session := engine.NewSession()

	//5.2.延迟释放
	defer session.Close()

	//5.3.开启事务
	session.Begin()

	//5.4.添加错误处理
	defer func() {
		err := recover()
		if err != nil {
			//回滚
			fmt.Println(err)
			fmt.Println("Rollback")
			session.Rollback()
		} else {
			session.Commit()
		}
	}()

	//6.事务内操作
	//6.1.插入数据
	user1 := User{Id: 10007, Name: "peixiaoze", Age: 18, Passwd: "12312312"}
	if _, err := session.Insert(&user1); err != nil {
		panic(err)
	}

	//6.2.修改数据
	user2 := User{Name: "peixiaoze222", Age: 3}

	if _, err := session.Where("id = 10002").Update(&user2); err != nil {
		panic(err)
	}
}
