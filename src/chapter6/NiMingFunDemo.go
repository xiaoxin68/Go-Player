package main

import "fmt"

//匿名函数：只能调用一次
func main() {
	//在定义匿名函数的时候直接调用
	sum, sub := func(num1 int, num2 int) (int, int) {
		return num1 + num2, num1 - num2
	}(10, 20)
	fmt.Println("sum=", sum, ",sub=", sub) //sum= 30 ,sub= -10
	
	//将匿名函数赋值给一个变量（函数变量）
	a := func(num1 int, num2 int) (int) {
		return num1 + num2
	}
	res := a(10,10)
	fmt.Println(res) //20
	
	//调用全局匿名函数
	res2 := Fun1(11,22)
	fmt.Println(res2)
}

//全局匿名函数
var (
	Fun1 = func(num1 int, num2 int) (int) {
		return num1 * num2
	}
)
