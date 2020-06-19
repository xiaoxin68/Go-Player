package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	var n1 = 100
	fmt.Printf("n1的类型是%T,n1占用的字节数是：%d \n",n1,unsafe.Sizeof(n1))
	
	var n2 int64 = 345
	fmt.Printf("n2的类型是%T,n2占用的字节数是：%d",n2,unsafe.Sizeof(n2))
}
