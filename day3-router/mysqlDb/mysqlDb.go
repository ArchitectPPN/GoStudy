package mysqlDb

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局对象db
var db *sql.DB

// InitDB 定义一个初始化数据库的函数
func InitDB() (db *sql.DB) {
	// DSN:Data Source Name
	dsn := "crm_dev_ddl:tPG66zbqj38sQFn9@tcp(rm-uf67h6texeyf9f325.mysql.rds.aliyuncs.com:3306)/crm_dev?charset=utf8"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("sql.Open Error: ", err)
		return
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		fmt.Println("sql.Ping Error: ", err)
		return
	}
	return
}

type user struct {
	id   int
	name string
}

// QueryMultiRowDemo 查询多条数据示例
func QueryMultiRowDemo() (userList map[int]user) {
	sqlStr := "select id, user_name from users where id > ? and id < ?"
	rows, err := db.Query(sqlStr, 0, 20)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	userList = make(map[int]user, 20)

	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		userList[u.id] = u
	}

	return userList
}
