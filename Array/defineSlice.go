package main

import "fmt"

func main()  {

	// 声明一个int类型， 长度和容量为 0 的 nil 切片
	var sliceInt2 []int

	var sliceOnlyLength [5]int //只指定长度，元素初始化为默认值0
	var sliceOnlyLength2 = [5]int{1,2,3,4} //指定长度为5， 但是初始化数据为4个， 剩下的一个会被默认设置为0
	fmt.Println(sliceOnlyLength)
	fmt.Println(sliceOnlyLength2)

	var sliceInt = []int{1,2,3,4,5,6}
	fmt.Println(sliceInt)
	fmt.Println(sliceInt2)

	// 向切片添加一个int值
	sliceInt = append(sliceInt, 90)
	fmt.Println(sliceInt)

	// 切片的声明
	var sliceA []string

	// 切片的赋值, 添加元素
	sliceA = append(sliceA, "string1")
	sliceA = append(sliceA, "string2")

	// 循环输出
	for _, val := range sliceA {
		fmt.Println(val)
	}

	// 使用make声明
	var sliceMake = make([]int, 3, 20)
	fmt.Println(sliceMake)

	/*
	 * 这是一个切片的声明：即声明一个没有长度的数组
	 */
	// 数组未创建
	// 方法1：直接初始化
	var s []int //声明一个长度和容量为 0 的 nil 切片
	var s1 = []int{1,2,3,4,5} // 同时创建一个长度为5的数组
	fmt.Println(s, s1)

	// 方法2：用make()函数来创建切片：var 变量名 = make([]变量类型,长度,容量)
	var s2 = make([]int, 0, 5)
	// 数组已创建
	// 切分数组：var 变量名 []变量类型 = arr[low, high]，low和high为数组的索引。
	var arr = [5]int{1,2,3,4,5}
	var slice []int = arr[1:4] // [2,3,4]
	fmt.Println(s, s2, slice)
}
