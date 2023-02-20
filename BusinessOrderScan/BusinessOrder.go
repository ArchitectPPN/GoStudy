package main

import (
	"BusinessOrderScan/BusinessOrderGoProductionScan"
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

func InitDb() (err error) {
	// 数据库连接
	if env == "local" {
		dsn = "root:@tcp(127.0.0.1:3306)/tools"
	} else if env == "dev" {
		dsn = "crm_dev:YKdYXQfhjjzDr281@tcp(rm-uf67h6texeyf9f325.mysql.rds.aliyuncs.com:3306)/crm_dev"
	} else if env == "pro" {
		dsn = "crm_pro_read:zXkRr9Jo7JpVCCkT@tcp(192.168.35.118:3306)/crm"
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

func main() {
	env = "dev"

	err := InitDb()
	if err != nil {
		fmt.Println("数据库连接失败~")
		return
	}

	BusinessOrderGoProductionScan.Run(dbObj, 0)

	//var caseInfo BusinessOrderGoProductionScan.CaseInfo
	//caseInfo = BusinessOrderGoProductionScan.GetCaseInfoByCaseId(dbObj, "100f9f1a-45ea-ae1b-f534-57620ea467a1")
	//fmt.Println(caseInfo.TeenagerInfo.String)

}
