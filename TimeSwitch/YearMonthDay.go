package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now().Format("20060102")

	fmt.Printf("%T \n", today)
	fmt.Println(today)

	//fmt.Println("年月日", time.Now().Year(), time.Now().Format("01"), time.Now().Day(), "|", time.Now().YearDay())
}
