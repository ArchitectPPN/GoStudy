package main

import (
	"fmt"
	"strings"
)

func main() {
	var mapTemp = map[string]string {
		"userName" : "Architect",
		"userAge" : "21",
		"userSex" : "male",
	}

	var echoJsonStr string
	echoJsonStr = echoJson(mapTemp)
	fmt.Println(echoJsonStr)
}

func echoJson(appInfoMap map[string]string) string {
	var jsonString string

	jsonString += "{"
	// map的值遍历
	for mapVal, mapIndex := range appInfoMap {
		jsonString += "\"" + mapIndex + "\"" + ":" + "\"" + mapVal + "\","
	}

	// 去除字符
	jsonString = strings.TrimRight(jsonString, ",")
	jsonString += "}"

	return jsonString
}