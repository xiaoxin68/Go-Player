package main

import "fmt"

func main() {
	num := 10

	fmt.Printf("num的地址为:%p\n",&num)
	fmt.Printf("num存储的内容为%d\n",num)

	fmt.Println("====================")

	var ptr1 *int
	ptr1 = &num
	fmt.Printf("ptr1的地址为:%p\n",&ptr1)
	fmt.Printf("ptr1的存储的地址为:%v\n",ptr1)
	fmt.Printf("ptr1的存储的地址指向的内容为%v\n",*ptr1)

	fmt.Println("====================")

	var ptr2 **int
	ptr2 = &ptr1
	fmt.Printf("ptr2的地址为:%p\n",&ptr2)
	fmt.Printf("ptr2的存储的地址为:%v\n",ptr2)
	fmt.Printf("ptr2的存储的地址指向的地址为%v\n",*ptr2)
	fmt.Printf("ptr2的存储的地址指向的内容为%v\n",**ptr2)
}
