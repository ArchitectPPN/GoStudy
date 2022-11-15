package main

import (
	"database/sql"
	"encoding/json"
	"example/queryCaseCode"
	"fmt"
	"gee"
	"net/http"
)

/*
(1)
$ curl -i http://localhost:9999/
HTTP/1.1 200 OK
Date: Mon, 12 Aug 2019 16:52:52 GMT
Content-Length: 18
Content-Type: text/html; charset=utf-8
<h1>Hello Gee</h1>

(2)
$ curl "http://localhost:9999/hello?name=geektutu"
hello geektutu, you're at /hello

(3)
$ curl "http://localhost:9999/hello/geektutu"
hello geektutu, you're at /hello/geektutu

(4)
$ curl "http://localhost:9999/assets/css/geektutu.css"
{"filepath":"css/geektutu.css"}

(5)
$ curl "http://localhost:9999/xxx"
404 NOT FOUND: /xxx
*/

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/getCaseInfoByCaseCode", func(c *gee.Context) {
		caseInfo := queryCaseCode.QueryCaseCode(c.Query("caseCode"))

		jsons, errs := json.Marshal(caseInfo) //转换成JSON返回的是byte[]
		if errs != nil {
			fmt.Println(errs.Error())
		}

		fmt.Println(string(jsons))
		c.String(http.StatusOK, string(jsons))
	})

	r.GET("/hello/:name", func(c *gee.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}

type user struct {
	id   int
	name string
}

// QueryRowDemo 查询单条数据示例
func QueryRowDemo(db *sql.DB) {
	sqlStr := "select id, user_name from users where id=?"
	var u user
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d user_name:%s \n", u.id, u.name)
}
