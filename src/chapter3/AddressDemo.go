package main

import "fmt"

func main() {
	var i int = 10
	fmt.Println("i的内存地址是：", &i) //i的内存地址是： 0xc00008a008

	//指针
	var ptr *int = &i
	fmt.Println("ptr是:", ptr)      //ptr是: 0xc00008a008
	fmt.Println("ptr的地址是:", &ptr)  //ptr的地址是: 0xc000094008
	fmt.Println("ptr指向的值是:", *ptr) //ptr指向的值是: 10

	*ptr++ //修改指针所指向的值后，i的内容也变了
	fmt.Println("ptr指向的值是:", *ptr)
	fmt.Println("i的值是：", i)

	i++ //修改i的值后，指针所指向的内容也变了
	fmt.Println("ptr指向的值是:", *ptr)
	fmt.Println("i的值是：", i)
}
