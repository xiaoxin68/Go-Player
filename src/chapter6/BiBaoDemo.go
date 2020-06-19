package main

import "fmt"

func main() {
	f := addNum()
	fmt.Println(f(11)) //21
	fmt.Println(f(12)) //33
	
	sf := addStr()
	fmt.Println(sf(" world ")) //hello world
	fmt.Println(sf(" zxx "))   //hello world  zxx
}

//注意：函数和它引用到的变量一起构成一个闭包
func addNum() func(int) int {
	var n int = 10 //只初始化一次
	return func(i int) int {
		n += i
		return n
	}
}

func addStr() func(string) string {
	var str string = "hello"
	return func(s string) string {
		str += s
		return str
	}
}
