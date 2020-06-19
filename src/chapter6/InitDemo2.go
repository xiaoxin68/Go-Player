package main

import (
	"fmt"
	"./utils"
)

//init函数的主要目的是完成初始化工作
var slary float64 = initSlary()
var address string = initAddress()

func init()  {
	fmt.Println("main - Init")
	slary = 18000.89
	address = "武汉"
}

func initSlary() float64 {
	fmt.Println("salary-init")
	return 10000.10
}
func initAddress() string {
	fmt.Println("address-init")
	return "随州"
}

func main()  {
	fmt.Println("main...")
	fmt.Printf("姓名：%s,年龄：%d，住址：%s，薪水：%f\n",utils.Name,utils.Age,address,slary)
}

//结果：
//age-init
//name-init
//util-init
//salary-init
//address-init
//main - Init
//main...
//姓名：星星,年龄：18，住址：武汉，薪水：18000.890000

//初始化顺序：
//1、util中的全局变量
//2、util中的init方法
//3、main中的全局变量
//4、main中的init方法
//5、main函数
