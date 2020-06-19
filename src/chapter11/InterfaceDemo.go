package main

import "fmt"

type AInterface interface {
	testA()
}
type BInterface interface {
	testB()
}
//一个接口可以继承其他接口。如果实现这个接口，则必须实现所有的方法
type CInterface interface {
	AInterface
	BInterface
	testC()
}
type Stu struct {}

func (stu Stu)testA()  {
	fmt.Println("testA")
}
func (stu Stu)testB()  {
	fmt.Println("testB")
}
func (stu Stu)testC()  {
	fmt.Println("testC")
}

//空接口：可以把任何一个变量赋值给空接口
type BlankInterface interface {}
func main() {
	stu := Stu{}
	var a AInterface = stu
	var b BInterface = stu
	var c CInterface = stu
	a.testA()
	b.testB()
	
	c.testA()
	c.testB()
	c.testC()
	
	//可以把任何一个变量赋给空接口
	var blank BlankInterface = stu
	blank = 10
	fmt.Println(blank)
	
	var blank2 interface{}
	blank2 = "hello"
	fmt.Println(blank2)
}
