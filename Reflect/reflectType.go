package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)

	if v.String() == "int" {
		fmt.Printf("type: %v value: %d \n", v, x)
	} else {
		fmt.Printf("type: %v value: %f \n", v, x)
	}
}

func main() {
	var a float32 = 3.141596
	reflectType(a)

	var b int = 2
	reflectType(b)
}