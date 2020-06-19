package main

import "fmt"
//结构体可以相互转化，但是要求它们的字段和方法全都一致
type A struct {
	Name string
}
type B struct {
	Name string
}

//给结构体取别名，go会认为是新的数据类型。但相互间可以强转
type AA A   //给结构体A去别名AA
func main() {
	var a A = A{"Zhangsan"}
	var b B
	b = B(a)
	fmt.Println(a, b)
	
	var aa AA
	aa = AA(a)//这里需要强转
	fmt.Println(aa)
}
