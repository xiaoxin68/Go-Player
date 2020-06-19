package main

import "fmt"

type P struct {
	Name string
	Age int
}
//给People类型绑定唯一方法
func (obj P) test()  {
	fmt.Println("test() name=",obj.Name)
}
func (obj P)getSum(num1,num2 int)int  {
	return num1+num2
}

func main() {
	var obj P  = P{
		Name: "张三",
		Age:  10,
	}
	obj.test()
	sum := obj.getSum(10, 20)
	fmt.Println(sum)
}
