package main

import "fmt"

type emptyInterface interface {
}

func main() {
	var emptyInterface emptyInterface
	fmt.Println(emptyInterface)

	emptyInterface = 18
	fmt.Println(emptyInterface)
	emptyInterface = "Hello"
	fmt.Println(emptyInterface)
	emptyInterface = false
	fmt.Println(emptyInterface)

	var mapTest = make(map[string]interface{}, 16)
	mapTest["name"] = "娜扎"
	mapTest["age"] = 16
	mapTest["hobby"] = []string{"藍球", "足球"}

	ret, isArray := mapTest["hobby"].([]string)

	if isArray {
		fmt.Println("You are right")
	} else {
		fmt.Println("You are error", ret)
		fmt.Printf("%T \n", mapTest["hobby"])
	}

	fmt.Println(mapTest)
}
