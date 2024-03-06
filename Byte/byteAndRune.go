package main

import "fmt"

func main() {
	s1 := "big"
	byteS1 := []byte(s1)
	// 修改b -> p
	byteS1[0] = 'p'
	fmt.Printf("%v %T %s %s \n", byteS1, byteS1, byteS1, s1)
	s2 := "我是一丹 s"
	rune2 := []rune(s2)
	rune2[2] = '二'
	rune2[3] = '蛋'
	rune2[5] = 'd'

	fmt.Printf("%s %v %s %T", s2, rune2, string(rune2), rune2)

	fmt.Println()

	traversalString()
}

func traversalString() {
	s := "hello 山河"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c) ", s[i], s[i])
	}

	fmt.Println()

	for _, r := range s {
		fmt.Printf("%v(%c) ", r, r)
	}
}
