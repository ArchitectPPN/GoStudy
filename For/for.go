package main

import "fmt"

func main() {
	//baseFor()
	//
	//baseFor2()
	//
	//sumCount()
	//
	//fmt.Println("换行")
	//
	//cutBamboo()
	//
	//fmt.Println("换行")

	traverseString()
}

func baseFor() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func baseFor2() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d", i)
	}
	fmt.Println('1')
}

func sumCount() {
	res := 0
	i := 1
	for i <= 100 {
		if i%3 == 0 {
			res += i
			fmt.Print(i)
			if i < 99 {
				fmt.Print("+")
			} else {
				fmt.Printf("=%d", res)
			}
		}
		i++
	}
}

func cutBamboo() {
	count := 0

	for i := 32.0; i >= 4; i -= 1.5 {
		count++
	}

	fmt.Print(count)
}

func traverseString() {
	str := "123abcdefg"
	for i, value := range str {
		fmt.Printf("第%d位的字符是: %c\n", i+1, value)
	}
}
