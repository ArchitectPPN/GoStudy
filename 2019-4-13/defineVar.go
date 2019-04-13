package main

import "fmt"

var (
	aInteger int
	bString  string
	cInt     int32
	dArray   []int
	eBool    bool
	fFunc    func() string
	l        = 200
)

func main() {
	m := 1
	n := 2
	m, n, q := 4, 5, 6

	aInteger = 190
	bString = "hhh"
	cInt = 2000
	// 交换两个变量的值
	m, n = n, m

	fmt.Printf("%T %v \n %T %v \n", aInteger, aInteger, bString, bString)
	fmt.Printf("%T %v \n %T %v \n", cInt, cInt, dArray, dArray)
	fmt.Printf("%T %v \n %T %v \n", eBool, eBool, fFunc, 1)
	fmt.Printf("%T %v \n %T %v \n", l, l, 12, 1)
	fmt.Printf("m= %T %v \n ", m, m)
	fmt.Printf("n= %T %v \n ", n, n)
	fmt.Printf("%T %v \n ", q, q)
}
