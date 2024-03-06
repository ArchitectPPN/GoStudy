package main

import "fmt"

func main()  {
	sayHello := func() {
		fmt.Println("Say Hello!")
	}

	sayHello()

	func(number1 int, number2 int) {
		number3 := number1 + number2
		fmt.Printf("%d", number3)
	}(1, 6)
}