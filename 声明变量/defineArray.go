package main

import "fmt"

func main(){

	var array1 [10] int

	var i, j int

	for i = 0; i < 10; i++ {
		array1[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Element [%d] = %d\n", j, array1[j])
	}

	var k int
	var arrayMap = [2]int{1,2}

	for k = 0; k < 2; k++ {
		fmt.Printf("Element [%d] = %d", k, arrayMap[k])
	} 

}
