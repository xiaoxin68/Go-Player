package main

import (
	util "../utils"  //支持取别名
	"fmt"
)

func main()  {
	var num1,num2 float64
	var op byte
	fmt.Println("请输入两个操作数和一个操作符，用逗号隔开：")
	fmt.Scanf("%f %f %c", &num1, &num2, &op)
	res := util.Cacullate(num1,num2,op)  //调用的是util包中的Cacullate
	fmt.Printf("%f%c%f=%f\n",num1,op,num2,res)
	fmt.Println(Cacullate(num1,num2,op))  //调用的是main方法中的Cacullate
	fmt.Println(cacullate(num1,num2,op)) //调用的是main方法中的cacullate
}
func Cacullate(num1 float64, num2 float64, op byte) float64 {
	return 0.1
}
func cacullate(num1 float64, num2 float64, op byte) float64 {
	return 0.2
}
