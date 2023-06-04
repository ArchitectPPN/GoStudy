package main

import (
	"context"
	"fmt"
	"github.com/blinkbean/dingtalk"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

type test struct {
	Val int
}

func main() {
	var tt *test

	var ttt test
	ttt.Val = 12

	tt = &ttt

	fmt.Println(tt)

	//fmt.Println(time.Now().Format("20060102"))
	//
	//// CreateFile
	//fileName := "A:\\tempFile\\createFile.txt"
	//// 创建文件
	//createCustomFile(fileName)
	//// 字符串写入文件中
	//writeStringToFile(fileName, "第一行文字~~~ \n")
	//
	//// 写入第二行文件
	//writeStringToFile(fileName, "第二行文字~~~ \n")
	//
	//// 连接redis
	//connectRds()
	//
	//// 发送钉钉
	//dingTalk()
}

// 创建文件
func createCustomFile(fileName string) {
	_, err := os.Create(fileName)
	if err != nil {
		fmt.Println("创建文件失败:", err)

		return
	}

	fmt.Println("创建文件成功！")
}

// 将字符串写入文件
func writeStringToFile(fileName, writeContent string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0200)
	if err != nil {
		fmt.Println("打开文件错误:", err)
		return
	}

	// 向文件中写入文本
	writeLength, err := file.WriteString(writeContent)
	if err != nil {
		fmt.Println("向文件写入字符时发生错误:", err)
		return
	}

	fmt.Println("写入了", writeLength, "字节内容")
	err = file.Close()
	if err != nil {
		fmt.Println("文件关闭出错:", err)
		return
	}
}

// 链接Rds
func connectRds() {
	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	rdb.Set(ctx, "newRds", "1111", 1000*time.Second)

	fmt.Println(rdb)
}

// 钉钉提醒
func dingTalk() {
	var token, secret string
	token = "30e727816845b120387412da50d164380802748d947cdbf6d5d948ec05901deb"
	secret = "SECaf679071ac68e5a4cc38a0d2f3c0e1bef91e656a41307f51526291c878babe47"

	cli := dingtalk.InitDingTalkWithSecret(token, secret)
	err := cli.SendTextMessage("我是测试Msg~")
	if err != nil {
		return
	}
}
