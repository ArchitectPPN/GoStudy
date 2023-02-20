package BusinessGoProductionScan

import "fmt"

type EmptyExmple struct {
	name string
}

var Tmp EmptyExmple

func New() {
	Tmp = EmptyExmple{"iii"}
	fmt.Println("success~")
}

func Two() {
	fmt.Println(Tmp.GetName())
}

func (r EmptyExmple) GetName() string {
	return r.name
}
