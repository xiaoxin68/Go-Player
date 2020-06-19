package main

import "fmt"

func main() {
	res := sumFunc(10, 20)
	fmt.Println("ok4,res = ", res) //4、ok4,res =  32
}

//defer的作用：在函数执行完毕后及时释放资源
func sumFunc(num1 int, num2 int) int {
	//当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的栈（defer栈）
	//当函数执行完毕后，再从defer栈，按照先入后出方式出栈执行
	//defer将语句拷入栈中，也会把值拷入栈中
	defer fmt.Println("ok1,num1 = ", num1) //3、ok1,num1 =  10
	defer fmt.Println("ok2,num2 = ", num2) //2、ok2,num2 =  20
	
	num1++
	num2++
	res := num1 + num2
	fmt.Println("ok3,res = ", res) //1、k3,res =  32
	return res
}

//输出结果：
//ok3,res =  32
//ok2,num2 =  20
//ok1,num1 =  10
//ok4,res =  32
