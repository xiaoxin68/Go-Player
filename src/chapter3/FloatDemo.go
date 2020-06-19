package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var n1 = -67.23 //默认占8字节float64
	fmt.Printf("n1的类型是%T,n1占用的字节数是：%d \n", n1, unsafe.Sizeof(n1))
	
	var n2 float32 = 232.4
	fmt.Printf("n2的类型是%T,n2占用的字节数是：%d \n", n2, unsafe.Sizeof(n2))
	
	var n3 float64 = -232.4
	fmt.Printf("n3的类型是%T,n3占用的字节数是：%d \n", n3, unsafe.Sizeof(n3))
	
	var num1 float32 = -12.0000901            //会损失精度
	var num2 float64 = -12.0000901            //原封不动保存我们的值
	fmt.Println("num1=", num1, "num2=", num2) //num1= -12.00009 num2= -12.0000901
	
	num5, num6 := 5.123,-.456 //在正式开发中还是把0带上
	fmt.Println("num5=", num5, "num6=", num6)//num5= 5.123 num6= -0.456
	
	num7,num8 := 10e2,11E-2 //科学计数法，e大写小写均可
	fmt.Println("num7=", num7, "num8=", num8)//num7= 1000 num8= 0.11
	
}
