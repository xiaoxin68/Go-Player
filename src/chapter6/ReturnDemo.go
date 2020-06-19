package main

import "fmt"

func main() {
	sum, sub := getSumAndsub(10, 20)
	fmt.Println("sum=", sum, ",sub=", sub) //sum= 30 ,sub= -10
}

//支持对函数返回值命名
func getSumAndsub(num1 int, num2 int) (sum int, sub int) {
	sum = num1 + num2
	sub = num1 - num2
	return sum, sub
}
