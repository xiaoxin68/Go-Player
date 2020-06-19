package main

import "fmt"

//如果一个结构体引入的多个嵌套结构体具有相同的属性或方法，
//则在访问时必须明确指明结构体名字，否则编译报错
type AA struct {
	Name string
	age int
}
type BB struct{
	Name string
	Score float64
}

//注意：如果嵌套的结构体有名字就是组合关系，在访问的时候必须带上结构体的名字
type CC struct {
	AA
	b BB //[组合关系]
	//Name string
}
func main() {
	var c CC
	//c.Name = "nameC"
	//如果CC里面没有Name属性时，赋值的时候要写明类型
	c.AA.Name = "nameA"
	//注意：如果嵌套的结构体有名字就是组合关系，在访问的时候必须带上结构体的名字
	c.b.Name = "nameB"
	//访问的时候也要写明
	fmt.Println(c.AA.Name,c.BB.Name)
}
