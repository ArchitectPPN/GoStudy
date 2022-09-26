package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 操作对象
var dbObj *sql.DB

// 当前环境
var env string

// 数据库描述
var dsn string

type user struct {
	id       int
	age      int
	userName string
}

type taskInfo struct {
	taskId string
}

func main() {
	// 设置当前环境
	env = "local"

	// 初始化数据库连接
	err := initDb()
	if err != nil {
		// 数据库初始化失败
		fmt.Println("数据库初始化失败", err)
		return
	}

	fetchOneRow(1)
}

// 初始化数据库连接
func initDb() (err error) {
	// 数据库连接
	if env == "local" {
		dsn = "root:@tcp(127.0.0.1:3306)/tools"
	} else if env == "dev" {
		dsn = "crm_dev:YKdYXQfhjjzDr281@tcp(rm-uf67h6texeyf9f325fo.mysql.rds.aliyuncs.com:3306)/crm_dev"
	}

	// open
	dbObj, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// 尝试与数据库建立连接
	err = dbObj.Ping()
	if err != nil {
		return err
	}

	return nil
}

func fetchOneRow(userId int) {
	querySql := "SELECT * FROM user WHERE id = ?"

	var user user
	err := dbObj.QueryRow(querySql, userId).Scan(&user.id, &user.userName, &user.age)
	if err != nil {
		fmt.Println("查询出错：", err)
		return
	}

	fmt.Println("查询到的数据为： ", user)
}
