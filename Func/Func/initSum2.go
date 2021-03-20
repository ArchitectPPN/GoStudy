package main

import "fmt"

func initSum2(x ...int) int {
	fmt.Println(x)

	sum := 0

	for _, value := range x {
		sum += value
	}

	return sum
}

func initSum3(x int, y ...int) int {
	fmt.Println(x, y)

	sum := x

	for _, value := range y {
		sum += value
	}

	return sum
}

func main() {
	resSum := initSum2(20, 30, 40, 50)
	fmt.Println(resSum)

	resSum4 := initSum3(20, 30, 40, 50, 60)
	fmt.Println(resSum4)

	resSum5, resSum6 := initSum4(1, 2)
	fmt.Println(resSum5, resSum6)

	resSum7, resSum8 := initSum5(11, 12)
	fmt.Println(resSum7, resSum8)

	resSum9 := returnNil()
	fmt.Printf("%T, %v", resSum9, resSum9)
}

// Go 语言支持多返回值, 但是必须用() 将其包裹
func initSum4(x, y int) (int, int) {
	return x, y
}

// 返回值命名
func initSum5(x, y int) (x1, y1 int) {
	x1, y1 = x, y
	return
}

func returnNil() []int {
	return []int{}
}
