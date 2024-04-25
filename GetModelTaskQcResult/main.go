package main

import (
	"GetModelTaskQcResult/QcResult"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 操作对象
var dbObj *sql.DB

// 当前环境
var env string

// 数据库描述
var dsn string

// InitDb 初始化数据库连接
func InitDb() (err error) {
	// 数据库连接
	if env == "local" {
		dsn = "root:@tcp(127.0.0.1:3306)/tools"
	} else if env == "dev" {
		dsn = "crm_dev:YKdYXQfhjjzDr281@tcp(rm-uf67h6texeyf9f325.mysql.rds.aliyuncs.com:3306)/crm_dev"
	} else if env == "sit" {
		dsn = "crm_sit:UOGRxk4SVvt0Zncl@tcp(rm-uf6v4cfq4a000sr3w.mysql.rds.aliyuncs.com:3306)/crm_sit"
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
	env = "sit"
	err := InitDb()
	if err != nil {
		fmt.Println("数据库连接失败~")
		return
	}

	var taskInfo QcResult.TaskInfo

	// 获取任务单id信息
	taskInfo = QcResult.GetTaskInfoByTaskName(dbObj, "T20230221005485")
	// 获取任务单质检信息
	QcResult.GetTaskQcInfo(dbObj, &taskInfo)
	QcResult.GetTaskReceipt(dbObj, &taskInfo)

	//fmt.Printf("%v", taskInfo)

	// 单个口内扫质检结果
	modelQcResult := QcResult.GetExampleQcResult(taskInfo)

	QcResultStr, _ := json.MarshalIndent(modelQcResult, "", "\t")

	fmt.Printf("%s", QcResultStr)
}
