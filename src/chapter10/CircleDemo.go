package main

import "fmt"

type Circle struct {
	ridus float64
}
//传入结构体的指针
func (c *Circle)area()float64  {
	c.ridus = 2.5
	//return 3.14*(*c).ridus*(*c).ridus  //标准的写法，但是也可以写成如下形式【底层做了优化】
	return 3.14*c.ridus*c.ridus
}
func main() {
	c1 := Circle{3.2}
	fmt.Println((&c1).area()) //标准的调用方式
	//也可以简写成如下这种形式
	fmt.Println(c1.area())
	fmt.Println(c1.ridus)
}
