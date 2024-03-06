package main

import "fmt"

func main()  {
	student := map[string]int {
		"age":18,
		"height":80,
	}

	for key, value := range student {
		fmt.Printf("key: %v, value: %v", key, value)
		fmt.Println("---")
	}
}