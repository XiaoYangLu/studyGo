package dbSql

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
)

var Db *sql.DB

// 连接数据库
func Db_mysql() {
	//配置参数
	config := beego.AppConfig
	dbDriver := config.String("db_driver")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	fmt.Println(dbDriver, dbUser, dbPassword, dbIp, dbName)

	//连接数据库
	connUrl := dbUser + ":" + dbPassword + "@tcp(" + dbIp + ")/" + dbName + "?charset=utf8"
	db, err := sql.Open(dbDriver, connUrl)
	if err != nil {
		panic("数据库连接有误！")
	}

	Db = db
}
