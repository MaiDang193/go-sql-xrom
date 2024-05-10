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

	/*
		-----xorm数据查询-----
	*/

	//5.Query
	results0, err := engine.Query("select * from user")
	fmt.Println(results0)

	results1, err := engine.QueryString("select * from user")
	fmt.Println(results1)

	results2, err := engine.QueryInterface("select * from user")
	fmt.Println(results2)

	//6.Get   只能查一条
	user := User{}
	engine.Get(&user)
	fmt.Println(user)

	//6.1.指定条件查询
	user1 := User{Name: "peixiaoze"}
	engine.Where("name = ?", user1.Name).Asc("id").Get(&user1)
	fmt.Println(user1)

	//6.2.获取指定字段的值
	var passwd string
	engine.Table(&user).Where("name = ?", "peixiaoze").Cols("passwd").Get(&passwd)
	fmt.Println(passwd)

	//7.Find  查询多条记录
	var users []User
	engine.Where("passwd = 123456").And("age = 18").Limit(10, 0).Find(&users)
	fmt.Println(users)

	//8.Count   select count(*) from user
	user = User{Age: 18}
	counts, err := engine.Count(&user)
	fmt.Println("总记录数为：", counts)

	//9.Iterate 和 Rows
	//9.1.Iterate
	engine.Iterate(&User{Passwd: "123456"}, func(idx int, bean interface{}) error {
		user := bean.(*User)
		fmt.Println("user:", user)
		return nil
	})
	// SELECT * FROM user

	//9.2.Rows
	rows, err := engine.Rows(&User{Name: "peixiaoze"})
	defer rows.Close()
	userBean := new(User)
	// SELECT * FROM user
	for rows.Next() {
		rows.Scan(userBean)
		fmt.Println(userBean)
	}
}
