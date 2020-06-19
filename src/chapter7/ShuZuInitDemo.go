package main

import "fmt"

//初始化数组的几种方式
func main() {
	//方式一
	var numArr1 [3]int  = [3]int{1,2,3}
	fmt.Println("numArr1=",numArr1)
	
	//方式二
	var numArr2 = [3]int{5,6,7}
	fmt.Println("numArr2=",numArr2)
	
	//方式三
	var numArr3 = [...]int{8,9,10}
	fmt.Println("numArr3=",numArr3)
	
	//方式四
	var numArr4 = [...]int{1:34,2:45,0:12}
	fmt.Println("numArr4=",numArr4)
	
	//类型推导
	strArrs := [...]string{1: "zxx", 2: "cy", 0: "ys"}
	fmt.Println("strArrs=",strArrs)
}
