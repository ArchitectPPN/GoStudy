package main

import (
	"Byte/MySqlDb"
	"Byte/ScanReceiptNoStl"
	"fmt"
)

func main() {
	// 排除掉没有硅胶条形码的硅胶
	// 病例结束的跳过
	// 数字化精修任务单直接跳过
	db, mysqlConfig := MySqlDb.Init("prod")

	fmt.Println("开始扫描:")
	ScanReceiptNoStl.Handle(db, mysqlConfig)
	fmt.Println("扫描结束:")
}
