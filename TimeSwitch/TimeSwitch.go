package main

import (
	"fmt"
	"time"
)

var timeTemplates = []string{
	"2006-01-02 15:04:05", //常规类型
	"2006/01/02 15:04:05",
	"2006-01-02",
	"2006/01/02",
	"15:04:05",
}

func TimeStringToGoTime(tm string) time.Time {
	for i := range timeTemplates {
		t, err := time.ParseInLocation(timeTemplates[i], tm, time.Local)
		if nil == err && !t.IsZero() {
			return t
		}
	}

	return time.Time{}
}

func test() {
	var tms = []string{
		"2021-03-04 08:15:11",
		"2021/03/05 08:15:11",
		"2021-03-04",
		"2021/03/05",
		"08:15:11",
	}

	for _, v := range tms {
		fmt.Println("OrgTimeStr: ", v, "; Convert Result: ", TimeStringToGoTime(v))
	}
}

func compareTime() {
	var tms = []string{
		"2021-03-04 08:15:11",
		"2022-03-04 08:15:11",
		"2023-03-04 08:15:11",
		"2019-03-04 08:15:11",
		"2015-03-04 08:15:11",
	}

	var bigTime time.Time
	for _, timeValue := range tms {
		temp := TimeStringToGoTime(timeValue)
		if bigTime.Before(temp) {
			bigTime = temp
		}
	}

	fmt.Println(bigTime)
}

func main() {
	//test()

	compareTime()
}
