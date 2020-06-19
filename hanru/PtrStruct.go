package main

import "fmt"

//结构体指针
func main() {
	var n1 = SS{"张三",12}
	fmt.Printf("n1的类型是%T，n1的地址是%p，n1的内容是%v\n",n1,&n1,n1)

	var n2 *SS
	n2 = &n1
	fmt.Printf("n2的类型是%T，n2的地址是%p，n2的内容是%v,n2的空间指向的内容是%v\n",n2,&n2,n2,*n2)

	n2.name = "李四"  //(*n2).name = "李四" 的简写
	fmt.Println(n1)
	fmt.Println(*n2)
}

type SS struct {
	name string
	age int
}
