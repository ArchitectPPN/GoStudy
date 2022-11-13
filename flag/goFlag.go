package main

import (
	"flag"
	"fmt"
)

var name string

func main() {
	flag.StringVar(&name, "name", "value", "user name")
	var name1 = flag.String("name1", "everyone", "The greeting object.")
	flag.Parse()

	fmt.Println(name, *name1)
}
