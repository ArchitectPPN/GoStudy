package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type userPerson struct {
	id       int
	userName string
	age      int
}

// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:@tcp(127.0.0.1:3306)/tools"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func queryOneRow(userId int) {
	querySql := "SELECT * FROM user WHERE id = ?"
	var oneUser userPerson
	err := db.QueryRow(querySql, userId).Scan(&oneUser.id, &oneUser.userName, &oneUser.age)
	if err != nil {
		fmt.Println("发生错误, err", err)
		return
	}

	fmt.Println(oneUser)
}

func queryMultiRowDemo(userId int) {
	querySql := "SELECT * FROM user WHERE id >= ?"
	rows, err := db.Query(querySql, userId)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}

	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	for rows.Next() {
		var userMuil userPerson
		err := rows.Scan(&userMuil.id, &userMuil.userName, &userMuil.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}

		fmt.Printf("id:%d name:%s age:%d\n", userMuil.id, userMuil.userName, userMuil.age)
	}
}

func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}

	queryOneRow(1)
	queryMultiRowDemo(1)

	fmt.Println("Connect Success~")
}
