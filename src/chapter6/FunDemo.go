package main

import "fmt"

//函数也是一种数据类型
func main() {
	a := getSum
	//a的类型是：func(int, int) int,getSum的类型是：func(int, int) int
	fmt.Printf("a的类型是：%T,getSum的类型是：%T\n", a, getSum)
	sum := a(1, 2) //等价于sum := getSum(1, 2)
	fmt.Println(sum)
	
	res := myFun(a, 10, 20)
	res2 := myFun(getSum, 20, 40)
	fmt.Printf("res=%d,res2=%d\n", res, res2)//res=30,res2=60
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

//既然函数也是一种类型，那么函数也可以作为一种形参
func myFun(mySum func(int, int) int, num1 int, num2 int) int {
	return mySum(num1, num2)
}
