package main

import (
	"fmt"
	"math"
)

type myFun func(float64) string
func main(){
	// 定义一个切片, 对其中的数据进行求平方根和求平方运算
	arr := []float64{1, 9, 16, 25, 30}

	result := FilterSlice(arr, func(val float64) string {
		val = math.Sqrt(val)
		return fmt.Sprintf("%.2f", val)
	})
	fmt.Print(result)

	result = FilterSlice(arr, func(val float64) string {
		val = math.Pow(val, 2)
		return fmt.Sprintf("%.0f", val)
	})
	fmt.Print(result)
}

// 遍历切片, 对其中每个元素进行运算处理
func FilterSlice(arr []float64, f myFun) []string {
	var result []string
	for _, value := range arr {
		result = append(result, f(value))
	}

	return result
}




