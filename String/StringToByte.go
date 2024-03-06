package main

import "fmt"

func main() {
	var userName string
	userName = "ArchitectPPN"
	var ch byte = 'A'

	b := []byte(userName)

	fmt.Printf("%s \n", userName)

	fmt.Println(b, ch)
}
