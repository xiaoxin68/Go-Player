package main

import "fmt"

//range 返回的是每个元素的副本，而不是直接返回对该元素的引用，如下所示。
func main() {
	// 创建一个整型切片，并赋值
	slice := []int{10, 20, 30, 40}
	// 迭代每个元素，并显示值和地址
	for index, value := range slice {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &slice[index])
	}
}
//Value: 10 Value-Addr: C00000A0C0 ElemAddr: C00000E340
//Value: 20 Value-Addr: C00000A0C0 ElemAddr: C00000E348
//Value: 30 Value-Addr: C00000A0C0 ElemAddr: C00000E350
//Value: 40 Value-Addr: C00000A0C0 ElemAddr: C00000E358
