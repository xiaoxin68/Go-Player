package main

import "fmt"

func main() {
	var num1 int8 //有符号整数[int8 int16 int32 int64]
	//测试int8的范围
	//num1 = 128  //int8[-128,127]    constant 128 overflows int8
	num1 = 127
	fmt.Println("num1=", num1)
	
	var num2 uint8 //无符号整数[uint8 uint16 uint32 unit64]
	//num2 = 256 //uint88[0,255]constant 256 overflows uint8
	num2 = 255
	fmt.Println("num2=", num2)
	
	var num3 int = -360  //64位系统占8个字节
	var num4 uint = 456    //54位系统占8个字节
	var num5 rune = 360  //与int32等价
	var num6 byte = 255   //与unit8等价[0,255]
	fmt.Println("num3=", num3,"\tnum4=",num4,"\tnum5=",num5,"\tnum6=",num6)
	
	var n1 = 100 //默认类型为int
	fmt.Printf("n1的类型为%T",n1)//Printf可用于格式化输出
}
