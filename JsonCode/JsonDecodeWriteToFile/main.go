package main

import (
	"encoding/json"
	"os"
)

type UserInfo struct {
	Name  string   `json:"name"`
	Age   int      `json:"age"`
	Hobby []string `json:"hobby"`
}

func main() {
	// 写入文件
	WriteToJson()

	// 更新json
	UpdateJson()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func WriteToJson() {
	user := &UserInfo{Name: "菜鸟", Age: 21, Hobby: []string{"看书", "编程", "GoLang", "Apex"}}
	data, err := json.MarshalIndent(user, "", "\t")
	checkErr(err)

	err = os.WriteFile("./JsonCode/JsonDecodeWriteToFile/data.json", data, 0755)
	checkErr(err)
}

func UpdateJson() {
	// 读取json数据
	data := make(map[string]interface{})
	file, err := os.ReadFile("./JsonCode/JsonDecodeWriteToFile/data.json")
	checkErr(err)
	err = json.Unmarshal([]byte(file), &data)
	checkErr(err)

	// 读取爱好数据
	hobbyList := make([]string, 0)
	for _, hobby := range data["hobby"].([]interface{}) {
		hobbyList = append(hobbyList, hobby.(string))
	}

	hobbyList = append(hobbyList, "玩游戏")
	hobbyList = append(hobbyList, "打Apex")
	data["hobby"] = hobbyList

	// 写入 data.json
	dataJson, err := json.MarshalIndent(data, "", "    ")
	checkErr(err)
	err = os.WriteFile("./JsonCode/JsonDecodeWriteToFile/data.json", dataJson, 0755)
	checkErr(err)
}
