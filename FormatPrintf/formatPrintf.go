package main

import "fmt"

type point struct {
	age, height int
	name        string
}

func main() {
	str := "我命由我不由天~"

	// 结构体
	p := point{25, 173, "阿牛"}

	fmt.Printf("%T, %v \n", str, str)
	fmt.Printf("%T, %v %t\n", p, p, p)

	fmt.Printf("%T %v \n", true, true)

	// 补充
	fmt.Printf("%-3d %-3d", 123223123, 1)

	// 按照格式化转化为特定的格式
	returnStr := fmt.Sprintf("%b", 123)

	fmt.Println(returnStr)

	// 输出字符的Unicode编码
	fmt.Printf("%U \n", '我')

	// 浮点数(保留2位)
	fmt.Printf("%.2f \n", 123.345)

	// 科学计数法
	fmt.Printf("%.1e \n", 123.123124124)

	// 字符串
	fmt.Printf("%s \n", "欢迎大家学校区块链")

}
