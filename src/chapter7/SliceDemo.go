package main

import "fmt"

func main() {
	//1、创建切片的第一种方式：通过数组
	var numArr [6]int = [...]int{1, 2, 3, 4, 5, 6}
	slice := numArr[1:3]
	fmt.Println("numArr=", numArr)
	fmt.Println("slice=", slice)
	fmt.Println("slice元素个数=", len(slice))
	fmt.Println("slice容量=", cap(slice)) //切片的容量可以动态变化:slice容量= 5
	
	//修改slice后，数组也变化了
	//slice中存储的内容：&numArr[1],len,cap
	slice[0] = 22
	fmt.Println("numArr=", numArr) //numArr= [1 22 3 4 5 6]
	fmt.Println("slice=", slice)   //slice= [22 3]
	
	//2、创建切片的第二种方式，通过make
	//make(类型，len，cap)：其中cap是可选参数，如果分配了，要求cap>=len
	var slice2 []int = make([]int, 5, 10)
	slice2[1] = 10
	slice2[3] = 30
	fmt.Println("slice2=", slice2)
	fmt.Println("slice2元素个数=", len(slice2))
	fmt.Println("slice2容量=", cap(slice2))
	
	//3、创建切片的第三种方式，定义一个切片，直接就指定数组
	var slice3 []int = []int{1, 2, 3}
	fmt.Println("slice3=", slice3)
	fmt.Println("slice3元素个数=", len(slice3)) //slice3元素个数= 3
	fmt.Println("slice3容量=", cap(slice3))   //slice3容量= 3
}
