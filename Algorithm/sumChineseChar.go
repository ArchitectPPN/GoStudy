package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 统计每个单子出现的次数
func main()  {
	s1 := "hello, 阿牛!"

	for _, value := range s1 {
		if unicode.Is(unicode.Han, value) {
			fmt.Printf("y")
		}
	}

	s2 := "how do you do"
	// 将字符串按照空格切开
	s3 := strings.Split(s2, " ")
	//
	m1 := make(map[string]int, 10)
	for _, value := range s3 {
		if _, ok := m1[value]; !ok {
			m1[value] = 1
		} else {
			m1[value]++
		}
	}

	fmt.Println(m1)
	for key, value := range m1 {
		fmt.Printf("%v, %v\n", key, value)
	}
}

