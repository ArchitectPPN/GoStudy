package main

import "fmt"

func main(){
	var score int = 99
	fmt.Scan("%d", score)
	getGreed(score)
}

func getGreed(score int){
	if score < 100 && score >= 90 {
		fmt.Printf("优秀")
	} else if score >= 70 && score <= 89 {
		fmt.Printf("良好")
	} else if score >= 60 && score <= 69 {
		fmt.Printf("及格")
	} else {
		fmt.Printf("不及格")
	}
}

