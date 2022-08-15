package gorm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"studyGo/frimes/config"
)

/*
 * 简单的连接mysq
 */
func DbMysqlPrivate() {
	//连接数据库
	//db, err := gorm.Open(config.DBOBJ, config.USER+"@("+config.DBIP+")/"+"?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", "root:117503@(127.0.0.1:3306)/gu?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接数据库错误：", err)
	}
	fmt.Println("连接成功")
	//关闭连接
	defer db.Close()

	//创建表
	db.AutoMigrate(&config.User{})
	//创建数据
	user := config.User{
		1,
		"张三",
		"男",
		"look Book",
	}
	db.Create(&user)
	//查询数据
	var u config.User
	db.First(&u)
	fmt.Println(u)
	//修改数据
	db.Model(&u).Update("Name", "李四")
	db.First(&u)
	fmt.Println(&u)

}
