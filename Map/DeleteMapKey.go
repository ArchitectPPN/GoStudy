package main

import "fmt"

func main()  {
	student := map[string]int{
		"小王":90,
		"小王1":80,
		"小王2":70,
		"小王3":60,
		"小王4":50,
	}

	for k, v := range student{
		if v < 80 {
			delete(student, k)
		}
	}

	fmt.Println(student)
}