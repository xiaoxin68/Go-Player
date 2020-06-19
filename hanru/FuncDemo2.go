package main

import "fmt"

//闭包：函数作为返回值
//闭包：一个外层函数中，有内层函数，该内层函数中，会操作外层函数的局部变量（外层函数中的参数，或者外层函数中直接定义的变量），
//并且该外层函数的返回值就是这个内层函数。这个内层函数和外层函数的局部变量，统称为闭包结构
//闭包中局部变量的生命周期会发生改变
func main() {
	rs1 := f2(0)        //指向了这个内层函数
	fmt.Println(rs1()) //1
	fmt.Println(rs1()) //2

	fmt.Println("================")

	//每次重新调用f2，就会重新生成局部变量i
	res2 := f2(1)
	fmt.Println(res2())
	fmt.Println(res2())

	fmt.Println("================")
	fmt.Println(rs1())
	fmt.Println(rs1())
}

//函数作为返回值：闭包
//虽然i是在函数内定义的局部变量，但是~~~
func f2(a int) func() int {
	i := 0 //每次调用内层函数的时候，这个i是共享的;这个局部变量的生命周期改变了
	return func() int {
		a++
		i++
		fmt.Println("a=",a,"i=",i)
		return a+i
	}
}
