package main

import "fmt"

func main()  {
	student := map[string]int{
		"age" : 89,
		"height" : 200,
	}

	value, exist := student["ages"]
	if exist {
		fmt.Printf("%v", value)
	} else {
		fmt.Printf("----")
	}
}
