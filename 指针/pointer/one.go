package pointer

import "fmt"

type oneDemo struct {
	Id   int
	name string
}

type Pointer struct {
	id   int
	name string
	demo oneDemo
}

func (c *Pointer) Test() {

}

func (c *Pointer) Init() {
	c.demo.Id = 1234
}

func (c *Pointer) EchoOneDemo() {
	one := new(oneDemo)

	fmt.Printf("%T", one)
}
