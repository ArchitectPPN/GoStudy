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

type stageLastModel struct {
	StageId      string
	CaseId       string
	MaterialType int
	ReceiptId    string
	Qualified    int
}

type caseInfo struct {
	CaseId       string
	CaseCode     string
	TeenagerInfo int
	CaseState    string
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

func getCaseInfo(w http.ResponseWriter, r *http.Request) {
	// 初始化数据库连接
	err := initDb()
	if err != nil {
		// 数据库初始化失败
		_, err := fmt.Fprintf(w, "数据库初始化失败, err: %v", err)
		if err != nil {
			return
		}
		return
	}

	caseIdStr := r.FormValue("caseId")

	var caseInfo caseInfo
	caseInfo = getCaseInfoByCaseId(caseIdStr)

	_, fErr := fmt.Fprintf(w, "{\"caseId\":\"%v\", \"CaseCode\":\"%v\", \"TeenagerInfo\":\"%v\", \"TeenagerInfoCn\":\"%v\", \"CaseState\":\"%v\", \"CaseStateCn\":\"%v\"}", caseInfo.CaseId, caseInfo.CaseCode, caseInfo.TeenagerInfo, caseUseA6OrAngelPole(caseInfo.TeenagerInfo), caseInfo.CaseState, caseStateMapping(caseInfo.CaseState))
	if fErr != nil {
		return
	}
}

func getCaseInfoByCaseId(caseId string) caseInfo {
	caseIdSql := "SELECT " +
		"ecc.id_c,ec.name,ecc.is_teenager_c,ecc.state_c " +
		"FROM ea_case ec " +
		"INNER JOIN ea_case_cstm ecc ON ecc.id_c = ec.id " +
		"WHERE ecc.id_c = ? "

	var caseInfo caseInfo
	err := dbObj.QueryRow(caseIdSql, caseId).Scan(&caseInfo.CaseId, &caseInfo.CaseCode, &caseInfo.TeenagerInfo, &caseInfo.CaseState)
	if err != nil {
		fmt.Println("查询出错：", err)
		return caseInfo
	}

	return caseInfo
}

func caseStateMapping(caseState string) (caseStateStr string) {
	/**
	#ecc.state_c 病例状态 1:资料处理中 4:3D设计中 5:3D设计待确认 6:加工中 7:已发货 8:暂停
	#                   9:结束 10:不收治 12:3D设计已确认 20:目标位设计中 21:目标位待确认
	#                   22:目标位已确认 30:产品待确认 31:产品已确认
	*/
	var caseStateMap = make(map[string]string, 10)
	caseStateMap["^1^"] = "资料处理中"
	caseStateMap["^4^"] = "3D设计中"
	caseStateMap["^5^"] = "3D设计待确认"
	caseStateMap["^6^"] = "加工中"
	caseStateMap["^7^"] = "已发货"
	caseStateMap["^8^"] = "暂停"
	caseStateMap["^9^"] = "结束"
	caseStateMap["^12^"] = "3D设计已确认"

	value, exist := caseStateMap[caseState]
	if exist {
		return value
	}

	return "未知状态"
}

func caseUseA6OrAngelPole(teenager int) (caseTeenager string) {
	var caseTeenagerMap = make(map[int]string, 3)

	caseTeenagerMap[0] = " "
	caseTeenagerMap[1] = "A6病例"
	caseTeenagerMap[2] = "下颌前导"

	value, exist := caseTeenagerMap[teenager]
	if exist {
		return value
	}

	return "不存在"
}

// 初始化数据库连接
func initDb() (err error) {
	// 数据库连接
	if env == "local" {
		dsn = "root:@tcp(127.0.0.1:3306)/tools"
	} else if env == "dev" {
		dsn = "crm_dev:YKdYXQfhjjzDr281@tcp(rm-uf67h6texeyf9f325.mysql.rds.aliyuncs.com:3306)/crm_dev"
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
	env = "dev"

	http.HandleFunc("/", apiIndex)
	http.HandleFunc("/getUserInfo", getUserInfoById)
	http.HandleFunc("/newUser", newUser)
	http.HandleFunc("/getCaseInfo", getCaseInfo)

	httpErr := http.ListenAndServe(":8090", nil)
	if httpErr != nil {
		fmt.Println("http server start err: ", httpErr)
		return
	}
}
