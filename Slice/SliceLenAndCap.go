package main

import "fmt"

func main()  {
	a := [5]int{1,2,3,4,5}
	s := a[3:5]

	fmt.Printf("s:%v, len(s):%d, cap(s):%v\n", s, len(s), cap(s))
	fmt.Println("")

	s2 := a[:] // 等同于 a[0:len(a)]
	fmt.Printf("s2:%v, len(s2):%d, cap(s2):%v\n", s2, len(s2), cap(s2))
	fmt.Println("")

	s3 := a[2:]  // 等同于 a[2:len(a)]
	fmt.Printf("s3:%v, len(s3):%d, cap(s3):%v\n", s3, len(s3), cap(s3))
	fmt.Println("")

	s4 := a[:3]  // 等同于 a[0:3]
	fmt.Printf("s4:%v, len(s4):%d, cap(s4):%v\n", s4, len(s4), cap(s4))
	fmt.Println("")
}