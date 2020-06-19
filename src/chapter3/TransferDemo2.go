package main

import "fmt"

func main() {
	var n1 int8 = 12
	var n2 int32
	var n3 int64
	
	n2 = int32(n1) + 20
	n3 = int64(n1) + 20
	
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)
	
	var num1 int32 = 20
	var num2 int8
	var num3 int8
	num2 = int8(num1) + 127 //编译通过，但是结果不正确，按溢出处理 num2= -109
	//num3 = int8(num1) + 128 //编译不通过
	fmt.Println("num1=", num1, "num2=", num2, "num3=", num3)
}
