package main

import (
	"fmt"
	"github.com/ArchitectPPN/golangHelper"
	"github.com/ArchitectPPN/golangHelper/mapInt"
	"github.com/ArchitectPPN/golangHelper/mapString"
)

func main() {
	var mapExample map[int]string
	mapExample = make(map[int]string)
	mapExample[0] = "12313"

	golangHelper.PrintShow()

	var mapExampleStr map[string]string
	mapExampleStr = make(map[string]string)
	mapExampleStr["test"] = "12313"

	exist := mapInt.CheckKeyExist(mapExample, 2)
	if exist {
		fmt.Println("key存在")
	} else {
		fmt.Println("key不存在")
	}

	exist = mapString.CheckKeyExist(mapExampleStr, "test")
	if exist {
		fmt.Println("key存在")
	} else {
		fmt.Println("key不存在")
	}
}
