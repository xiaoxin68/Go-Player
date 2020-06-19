package main

import "fmt"

//基本类型数据转换
func main() {
	var n1 int32 = 100
	
	var f1 float32 = float32(n1)
	fmt.Printf("%f\n", f1) //100.000000
	
	var n2 int8 = int8(n1)  //高精度->低精度  注意：变量n1本身数据类型并没有发生变化
	var n3 int64 = int64(n1)  //低精度->高精度
	fmt.Println("n2=",n2,"n3=",n3)
	
	var num1 int16 = 128
	num2 :=  int8(num1) //int8范围是[-128,127]会溢出
	fmt.Println("num2=",num2) //num2= -128  但是此处不报错，按溢出处理
	
}
