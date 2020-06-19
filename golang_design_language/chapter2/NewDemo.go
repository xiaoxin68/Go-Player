package main

import "fmt"

func main() {
	//每次调用new函数返回一个具有唯一地址的不同变量
	p := newInt1()
	q := newInt2()
	fmt.Println(p == q) //false
}

//newInt1和newInt2函数有同样的行为
func newInt1()*int  {
	return new(int)
}

func newInt2()*int  {
	var temp int
	return &temp
}