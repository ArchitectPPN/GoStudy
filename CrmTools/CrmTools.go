package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
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

type stageLastModel struct {
	StageId      string
	CaseId       string
	MaterialType int
	ReceiptId    string
	Qualified    int
}

func apiIndex(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "WelCome Api Index !\n")
	if err != nil {
		return
	}
}

func getUserInfoById(w http.ResponseWriter, r *http.Request) {
	// 初始化数据库连接
	err := initDb()
	if err != nil {
		// 数据库初始化失败
		fmt.Println("数据库初始化失败", err)
		return
	}

	userIdStr := r.FormValue("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return
	}

	userInfo := fetchOneRow(userId)

	_, fErr := fmt.Fprintf(w, "{\"userId\":\"%v\", \"userAge\":\"%v\", \"userName\":\"%v\"}", userInfo.id, userInfo.age, userInfo.userName)
	if fErr != nil {
		return
	}
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

func fetchOneRow(userId int) user {
	querySql := "SELECT * FROM user WHERE id = ?"

	var user user
	err := dbObj.QueryRow(querySql, userId).Scan(&user.id, &user.userName, &user.age)
	if err != nil {
		fmt.Println("查询出错：", err)
		return user
	}

	fmt.Println("查询到的数据为： ", user)

	return user
}

func insertNewUser(user user) {
	err := initDb()
	if err != nil {
		return
	}
	sqlStr := "INSERT INTO user(user_name, age) values (?, ?)"
	ret, err := dbObj.Exec(sqlStr, user.userName, user.age)
	if err != nil {
		fmt.Println("insert failed, err", err)
		return
	}

	theNewUserI, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Println("get last id failed! err:", err)
		return
	}

	fmt.Println("Insert success, the new user id ", theNewUserI)
}

type userData struct {
	UserAge  int    `json:"UserAge"`
	UserName string `json:"UserName"`
	UserId   int    `json:"UserId"`
}

func newUser(w http.ResponseWriter, r *http.Request) {
	var userData userData
	var user user
	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		err := r.Body.Close()
		if err != nil {
			return
		}
	}

	jsonUserDat, _ := json.Marshal(userData)
	fmt.Println("数据:", jsonUserDat)
	fmt.Println("数据：", userData)

	user.id = userData.UserId
	user.age = userData.UserAge
	user.userName = userData.UserName

	insertNewUser(user)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(userData)

	if err != nil {
		return
	}
}

func main() {
	// 设置当前环境
	env = "local"

	http.HandleFunc("/", apiIndex)
	http.HandleFunc("/getUserInfo", getUserInfoById)
	http.HandleFunc("/newUser", newUser)

	httpErr := http.ListenAndServe(":8090", nil)
	if httpErr != nil {
		fmt.Println("http server start err: ", httpErr)
		return
	}
}
