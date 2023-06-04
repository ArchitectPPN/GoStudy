package main

import (
	"fmt"
	"strconv"
)

func main() {

	aInt, err := strconv.ParseInt("1234", 10, 8)
	if err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Printf("数据类型: %T, 值: %v", aInt, aInt)
}
