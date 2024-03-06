package main

import "fmt"

var aa = "hello,go"

func main() {
	fmt.Printf("%08b", -29)
	fmt.Println(aa)

	aa = "wosho"
	fmt.Println(aa)
}
