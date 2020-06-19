package main

import "fmt"

func main() {
	//num1的类型为int，num1的值是100，num1的地址是0xc00000a0c0
	num1 := 100
	fmt.Printf("num1的类型为%T，num1的值是%v，num1的地址是%v\n",num1,num1,&num1)
	
	//new用来给值变量分配内存空间，返回的是指针类型
	//num2的类型为*int，num2的值是0xc00000a0d8，num2的地址是0xc000006030,num2指针指向的值是200
	num2 := new(int)  //num2位指针类型
	*num2 = 200
	fmt.Printf("num2的类型为%T，num2的值是%v，num2的地址是%v,num2指针指向的值是%v\n",num2,num2,&num2,*num2)
}