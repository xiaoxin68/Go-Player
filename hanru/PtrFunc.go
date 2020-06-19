package main

import "fmt"

func main() {
	//函数指针：一个指针，指向了一个函数的指针。因为go语言中，function默认看做一个指针，没有*
	var f1 func()[4]int
	f1 = fun11
	fmt.Printf("f1的类型为%T，地址%p，数值：%v\n",f1,&f1,f1)
	res := f1()
	fmt.Printf("res的类型为%T，地址%p，数值：%v\n",res,&res,res)

	fmt.Println("====================")

	//指针函数：
	var f2 func()*[4]int
	f2 = fun22
	fmt.Printf("f2的类型为%T，地址%p，数值：%v\n",f2,&f2,f2)
	res2 := f2()
	fmt.Printf("res2的类型为%T，地址%p，存储的地址：%v，存储的地址指向的内容%v\n",res2,&res2,res2,*res2)
}
//普通函数
func fun11() [4]int {
	return [...]int{1,2,3,4}
}

//指针函数
func fun22()*[4]int  {
	return &[...]int{1,2,3,4}
}
