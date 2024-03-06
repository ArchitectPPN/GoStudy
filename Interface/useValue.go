package main

import "fmt"

type dog struct {
}

func (d dog) say() {
	fmt.Println("汪汪~")
}

type cat struct {
}

func (c cat) say() {
	fmt.Println("喵喵~")
}

type sayers interface {
	say()
}

func da(arg sayers) {
	arg.say()
}

func main() {
	cat1 := cat{}
	da(cat1)

	dog1 := dog{}
	da(dog1)

	var s sayers
	s = cat1
	fmt.Println(s)
}
