package main

import (
	"encoding/json"
	"fmt"
)

type Seller struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	CountryCode string `json:"country_code,omitempty"` // 会忽略掉空值
}
type Seller2 struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	CountryCode string `json:"country_code"` // 输出空值
}
type Seller3 struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	CountryCode string `json:"-"` // 忽略字段
}

type Product struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Seller Seller `json:"seller"`
	Price  int    `json:"price"`
}

func main() {
	//From map
	fileCount := map[string]int{
		"cpp": 10,
		"js":  8,
		"php": 7,
	}

	bytes, _ := json.Marshal(fileCount)
	fmt.Printf("json: %s \n", string(bytes))

	// From Struct
	type Author struct {
		UserName string `json:"userName"`
		Age      int    `json:"age"`
		Sex      int    `json:"sex"`
	}
	type Book struct {
		Title  string `json:"title"`
		Author Author `json:"userInfo"`
		Year   int    `json:"year"`
		Month  int    `json:"month"`
		Day    int    `json:"day"`
	}

	book1 := Book{"你好，世界！", Author{"坐着", 20, 1}, 1995, 12, 12}
	structBytes, _ := json.MarshalIndent(book1, "", "\t")
	fmt.Printf("from struct to json: %s \n", structBytes)

	// 忽略JSON输出中的特定字段
	seller1 := Seller{Id: 1, Name: "CountryCode不为空", CountryCode: "US"}
	seller2 := Seller2{Id: 2, Name: "CountryCode为空", CountryCode: ""}
	seller3 := Seller3{Id: 2, Name: "CountryCode不会参与编码", CountryCode: "AAA"}
	seller4 := Seller2{Id: 2, Name: "CountryCode未赋值"}
	seller5 := Seller2{Id: 2, Name: "输出格式化json， prefix", CountryCode: "格式化输出"}

	ignoreEmpty, _ := json.Marshal(seller1)
	fmt.Printf("转换时忽略为空：%s \n", ignoreEmpty)

	normal, _ := json.Marshal(seller2)
	fmt.Printf("不忽略字段转换：%s \n", normal)

	normal2, _ := json.Marshal(seller4)
	fmt.Printf("CountryCode字段未赋值转换：%s \n", normal2)

	joinCant, _ := json.Marshal(seller3)
	fmt.Printf("不参与字段转换：%s \n", joinCant)

	formatJsonStr, _ := json.MarshalIndent(seller5, "Prefix", "\t")
	fmt.Printf("不参与字段转换：%s \n", formatJsonStr)
}
