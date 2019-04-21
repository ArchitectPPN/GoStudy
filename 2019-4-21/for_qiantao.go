package main

import "fmt"

func main() {
	//printRectangle()

	printRightTangle()

	printLeftTopTangle()
}

const (
	LENGTH = 9
)

// 打印矩形
func printRectangle() {
	fmt.Println("打印一个矩形")
	for i := 0; i < LENGTH; i++ {
		for j := 0; j < LENGTH; j++ {
			fmt.Print("❤")
		}
		fmt.Println("")
	}
}

// 打印左下直角三角形
func printRightTangle() {
	fmt.Println("打印一个左下角直角三角形")
	for i := 1; i < LENGTH; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("❤")
		}
		fmt.Println("")
	}
}

func printLeftTopTangle() {
	fmt.Println("打印左上角直角三角形")
	for i := 0; i < LENGTH; i++ {
		for j := LENGTH; j > i; j-- {
			fmt.Print("❤")
		}
		fmt.Println("")
	}
}
