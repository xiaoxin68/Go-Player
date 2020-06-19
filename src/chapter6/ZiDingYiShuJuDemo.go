package main

import "fmt"

func main() {
	//自定义数据类型
	type myInt int
	var i myInt = 10
	var j int
	j = int(i)
	fmt.Printf("i的类型是%T，i=%d\n", i, i) //i的类型是main.myInt，i=10
	fmt.Printf("j的类型是%T，j=%d\n", j, j)//j的类型是int，j=10
	
	res := add(funVar,10,20)
	fmt.Println(res)//30
}
func test(num1 int, num2 int) int {
	return num1 - num2
}
//自定义数据类型
type myTest func(int,int)int
func add( f myTest,num1 int,num2 int)int  {
	return f(num1,num2)
}
func funVar(num1 int,num2 int) int {
	return num1 + num2
}
