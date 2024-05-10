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

	//4.1.engine.Sync(new(User))创建User并完成同步
	err = engine.Sync(new(User))
	if err != nil {
		fmt.Println("创建同步失败！！！")
	}

	//5.增删改查
	//5.1.engine.Insert()插入对象，返回值：受影响的行数
	//5.1.1.插入单个对象
	user := User{Id: 10002, Name: "peixiaoze", Age: 18, Passwd: "123456"}
	n, _ := engine.Insert(&user)
	fmt.Println(n)
	if n >= 1 {
		fmt.Println("数据插入成功！！！")
	}

	//5.1.2.插入多个对象
	user1 := User{Id: 10003, Name: "peixiaoze", Age: 18, Passwd: "123456"}
	user2 := User{Id: 10004, Name: "peixiaoze", Age: 18, Passwd: "123456"}
	n, _ = engine.Insert(&user1, &user2)
	fmt.Println(n)
	if n >= 1 {
		fmt.Println("数据插入成功！！！")
	}

	//5.1.3.插入切片对象
	var users []User
	users = append(users, User{Id: 10005, Name: "peixiaoze", Age: 18, Passwd: "123456"})
	users = append(users, User{Id: 10006, Name: "peixiaoze", Age: 18, Passwd: "123456"})
	n, _ = engine.Insert(&users)

}
