package main

import "fmt"

func initSum(x, y int) int {
	return x + y
}

func main()  {
	fmt.Printf("%v\n", initSum(1, 20))

	sumRes := initSum(10, 20)
	fmt.Println("sumRes", sumRes)
}
