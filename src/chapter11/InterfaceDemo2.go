package main

import "fmt"

type AIter interface {
	Say()
}

type M struct {}
func (m M)Say() {//注意：这里的接受参数是结构体类型
	fmt.Println("m say...")
}

type N struct {}
func (n *N)Say()  { //注意：这里的接受参数是指针类型
	fmt.Println("n say...")
}
func main() {
	var m M = M{}
	var a AIter = m  //正确的写法
	a.Say()
	
	var a1 AIter = &N{}//正确的写法
	a1.Say()
}
