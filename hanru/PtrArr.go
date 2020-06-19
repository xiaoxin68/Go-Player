package main

import "fmt"

//数组指针和指针数组
func main() {
	//数组指针：存储数组的指针

	ptr1 := &[4]int{1, 2, 3, 4}

	fmt.Printf("ptr1的类型为:%T\n", ptr1)
	fmt.Printf("ptr1的地址为:%p\n", &ptr1)
	fmt.Printf("ptr1的存储的地址为:%v\n", ptr1)
	fmt.Printf("ptr1的存储的地址指向的内容为%v\n", *ptr1)

	fmt.Println("====================")

	arr := [4]int{1, 2, 3, 4}
	var ptr2 *[4]int
	ptr2 = &arr
	fmt.Printf("ptr2的类型为:%T\n", ptr2)
	fmt.Printf("ptr2的地址为:%p\n", &ptr2)
	fmt.Printf("ptr2的存储的地址为:%v\n", ptr2)
	fmt.Printf("ptr2的存储的地址指向的内容为%v\n", *ptr2)

	fmt.Println("====================")

	//数组指针：存储指针的数组
	a := 1
	b := 2
	c := 3
	d := 4
	var arr2 = [4]*int{&a, &b, &c, &d}
	fmt.Printf("arr2的类型为:%T\n", arr2)
	fmt.Printf("arr2的地址为:%p\n", &arr2)
	fmt.Printf("arr2的存储的地址指向的内容为%v\n", arr2)
}
