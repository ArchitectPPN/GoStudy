package main

import "fmt"

func main()  {
	student := map[string]int{
		"age":19,
		"height": 200,
	}

	fmt.Printf("%v", student)
	delete(student, "age")
	fmt.Printf("%v", student)
}
