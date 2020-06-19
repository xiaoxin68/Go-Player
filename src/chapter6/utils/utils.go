package utils  //只有main包才可被编译成可执行文件

import "fmt"

func Cacullate(num1 float64, num2 float64, op byte) float64 {
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
