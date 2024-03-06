package main

import "fmt"

func main()  {
	var mapping = [3][2]int{{1, 2}, {5, 6}, {3, 4}}

	for _, value := range mapping {
		for key, items := range value {
			fmt.Printf("index %d, Value: %d \n", key, items)
		}
	}
}
