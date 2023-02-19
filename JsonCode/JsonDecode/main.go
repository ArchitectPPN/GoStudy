package main

import (
	"encoding/json"
	"fmt"
)

type Window struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Title  string `json:"title"`
}

func main() {
	jsonInput := `{"width":500, "Height": 200, "Title":"我是windows窗口！"}`

	var window Window
	err := json.Unmarshal([]byte(jsonInput), &window)

	if err != nil {
		fmt.Println("json decode err!", err)
		return
	}

	fmt.Println(window)
	fmt.Printf("高：%d 宽：%d 标题：%s \n", window.Width, window.Height, window.Title)
}
