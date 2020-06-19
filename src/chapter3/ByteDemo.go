package main

import "fmt"

//go语言编码utf-8，苏破译基本上不会出现中文乱码的问题
func main() {
	//字符类型颖byte存储
	var b1 byte = 'a'
	b2 := '0'
	fmt.Println("b1=", b1, "b2=", b2) //b1= 97 b2= 48
	
	fmt.Printf("b1=%c,b2=%c \n", b1, b2) //b1=a,b2=0
	
	var b3 = '北'                           //var b3 int = '北'
	fmt.Printf("b3=%c,b3对应的码值为%d", b3, b3) //b3=北,b3对应的码值为21271
	
	//这里byte是存不进去的，得用int来存储
	//var b4  byte = '北'// constant 21271 overflows byte
	
	//允许使用转义字符来表示特殊字符
	var c1 = '\n'
	fmt.Printf("c1=%c,c1=%c \n", c1, c1) //,c1=
	
	//英文字符1个字节，汉字3个字节
	
	var c2 int = 21271                       //可以直接给字符变量赋值它的ASCII码值
	fmt.Printf("c2=%c,c2对应的码值为%d\n", c2, c2) //c2=北,c2对应的码值为21271
	
	var n1 = 10 + 'a'                               //10 + 97
	fmt.Printf("n1表示的字符是%c,n1的ASCII码是%d\n", n1, n1) //n1表示的字符是k,n1的ASCII码是107
	
}
