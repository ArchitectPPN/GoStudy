package main

import "fmt"

var arrNum = [5]int{1, 3, 5, 7, 8}

func main()  {

	target := 8
	num := 1
	temp := 0

	var res = [2][2]int{}

	// 检查是否在数组中
	temp = target - num
	if checkNumInArr(temp, arrNum) {

		res[0][0] = findIndex(temp, arrNum) - 1
		res[0][1] = temp

		res[1][0] = findIndex(num, arrNum) - 1
		res[1][1] = num

	} else {
		return
	}

	fmt.Println(res)
}

func checkNumInArr(x int, temp [5]int) bool  {
	for _, tempVal := range temp {
		if tempVal == x {
			return true
		}
	}

	return false
}

func findIndex(x int, temp [5]int) int  {
	for temKey, tempVal := range temp {
		if tempVal == x {
			return temKey + 1
		}
	}

	return 0
}