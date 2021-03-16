package main

import "fmt"

type student struct {
	name        string
	age, height int
}

func main() {
	var s1 student
	s1.age = 10
	s1.name = "小王八"
	s1.height = 180

	fmt.Println(s1.height)
}
