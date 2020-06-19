package main

import "fmt"

func main() {
	var slice1 = []int{1, 2, 3, 4}
	slice2 := make([]int, 10)
	//copy返回被复制的元素数量，它会是 len(src) 和 len(dst) 中较小的那个
	len := copy(slice2, slice1)    //注意：切片类型才可以拷贝
	fmt.Println("len=", len)       //len= 4
	fmt.Println("slice1=", slice1) //slice1= [1 2 3 4]
	fmt.Println("slice2=", slice2) //slice2= [1 2 3 4 0 0 0 0 0 0]
	
	//修改slice2，不会影响slice1：它们的数据空间是独立的
	slice2[2] = 22
	fmt.Println("slice1=", slice1) //slice1= [1 2 3 4]
	fmt.Println("slice2=", slice2) //slice2= [1 2 22 4 0 0 0 0 0 0]
	
	len2 := copy(slice1, slice2)
	fmt.Println("len2=", len2)     //len= 4
	fmt.Println("slice1=", slice1) //slice1= [1 2 22 4]
	fmt.Println("slice2=", slice2) //slice2= [1 2 22 4 0 0 0 0 0 0]
}
