package main

import "fmt"

type mover interface {
	move()
}

type sayer interface {
	say()
}

type person struct {
	name string
	age  int8
}

func (p *person) move() {
	fmt.Printf("%s running \n", p.name)
}

func (p *person) say() {
	fmt.Printf("%s say %d \n", p.name, p.age)
}

func main() {
	var m mover
	var s sayer
	p1 := person{
		name: "王子",
		age:  25,
	}

	p2 := &person{
		name: "pointer",
		age:  18,
	}

	m = p2
	m = &p1
	s = &p1
	m.move()
	s.say()
	fmt.Println(s)
}
