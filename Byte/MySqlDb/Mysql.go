package MySqlDb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func connectDb(dsn string) (dbObj *sql.DB) {
	// open
	dbObj, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("数据库连接失败, 原因: ", err)
		return
	}

	// 尝试与数据库建立连接
	err = dbObj.Ping()
	if err != nil {
		fmt.Println("与数据库建立连接失败, 原因: ", err)
		return dbObj
	}

	return
}

func Init(env string) (*sql.DB, *MySqlConfig) {
	// 初始化mysql
	mysqlConfig := NewMysqlConfig()

	switch env {
	case "dev":
		mysqlConfig.SetDevConfig()
	case "sit":
		mysqlConfig.SetSitConfig()
	case "prod":
		mysqlConfig.SetProdConfig()
	}

	dsn := mysqlConfig.GetMysqlDbConfig()

	return connectDb(dsn), mysqlConfig
}
