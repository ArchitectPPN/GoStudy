package main

import (
	"fmt"
	"time"
)

type Book struct {
	title      string
	totalPage  int16
	author     string
	publicTime string
}

func main() {
	var books = make(map[int]*Book, 10)
	var sanGuoYanYi = Book{"三国演义", 500, "罗贯中", "2019-01-09"}
	var hongLou = Book{"红楼梦", 200, "曹雪芹", "2022-08-09"}
	var shuiHu = Book{"水浒传", 400, "施耐庵", "2023-09-10"}
	var xiYouJi = Book{"西游记", 300, "吴承恩", "2029-10-11"}

	books[1] = &sanGuoYanYi
	books[2] = &hongLou
	books[3] = &shuiHu
	books[4] = &xiYouJi

	var lastBook *Book
	fmt.Println(lastBook)

	for _, mapValue := range books {
		if lastBook == nil {
			lastBook = mapValue
		} else {
			time1, err := time.Parse("3019-01-09 00:00:00", lastBook.publicTime)
			time2, err := time.Parse("3019-01-09 00:00:00", mapValue.publicTime)
			if err == nil && time2.After(time1) {
				lastBook = mapValue
			} else {
				fmt.Println("时间转换出错：", err)
			}
		}

		//fmt.Printf("第 %d 本书, 书名: %s, 总页数: %d, 作者： %s \n", mapKey, mapValue.title, mapValue.totalPage, mapValue.author)

	}

	fmt.Printf("最新书籍, %v \n", *lastBook)

	var pointerBook *Book
	pointerBook = &sanGuoYanYi
	fmt.Printf("title: %s, 总共：%d 页, 作者：%s", pointerBook.title, pointerBook.totalPage, pointerBook.author)
}
