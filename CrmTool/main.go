package main

import (
	_ "CrmTool/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	beego.Run()
}

// init 注册数据库orm
func init() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		beego.Info("注册MySql失败, err:", err)
		return
	}

	/**
	 *	Url:rm-uf6v4cfq4a000sr3w.mysql.rds.aliyuncs.com
	 *	A:crm_sit
	 *	P:UOGRxk4SVvt0Zncl
	 *  Dsn: username:password@tcp(host:port)/Dbname?charset=utf8&parseTime=True&loc=Local
	 *  Dsn: crm_sit:UOGRxk4SVvt0Zncl@tcp(rm-uf6v4cfq4a000sr3w.mysql.rds.aliyuncs.com:3306)/crm_sit?charset=utf8&parseTime=True&loc=Local
	 */
	//err = orm.RegisterDataBase("default", "mysql", "crm_sit:UOGRxk4SVvt0Zncl@tcp(rm-uf6v4cfq4a000sr3w.mysql.rds.aliyuncs.com:3306)/crm_sit?charset=utf8&parseTime=True&loc=Local")
	err = orm.RegisterDataBase("default", "mysql", "crm_dev:YKdYXQfhjjzDr281@tcp(rm-uf67h6texeyf9f325.mysql.rds.aliyuncs.com:3306)/crm_dev?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		beego.Info("注册默认数据库, err:", err)
		return
	}
}
