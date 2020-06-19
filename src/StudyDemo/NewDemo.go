package main

import "fmt"

//new() 函数可以创建一个对应类型的指针，创建过程会分配内存，被创建的指针指向默认值。
func main() {
	str := new(string)
	*str = "hello张三"
	*str += "李四"
	fmt.Println(*str)
}
