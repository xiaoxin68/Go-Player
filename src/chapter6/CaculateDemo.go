package main

import "fmt"

func main() {
	var num1,num2 float64
	var op byte
	fmt.Println("请输入两个操作数和一个操作符，用逗号隔开：")
	fmt.Scanf("%f %f %c", &num1, &num2, &op)
	res := cacullate(num1,num2,op)
	fmt.Printf("%f%c%f=%f",num1,op,num2,res)
}

func cacullate(num1 float64, num2 float64, op byte) float64 {
	var res float64
	switch op {
	case '+':
		res = num1 + num2
	case '-':
		res = num1 - num2
	case '*':
		res = num1 * num2
	case '/':
		if num2 != 0 {
			res = num1 / num2
		} else {
			res = 0
		}
	default:
		fmt.Println("输入错误")
	}
	return res
}
