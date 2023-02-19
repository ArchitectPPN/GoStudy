package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonInput := `{"apple":"10", "mango":"20", "grapes":"20"}`
	var fruitBasket map[string]string
	err := json.Unmarshal([]byte(jsonInput), &fruitBasket)
	if err != nil {
		fmt.Println("Json decode error!", err)
		return
	}

	fmt.Println(fruitBasket)
}
