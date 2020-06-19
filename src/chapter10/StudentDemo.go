package main

import "fmt"

type Student struct {
	Name string
	Age int
}
//如果一个类型实现了String方法，那么fmt.Println默认会调用这个方法进行输出
func (stu *Student)String()string  {
	str := fmt.Sprintf("[姓名=%v,年龄=%v]",stu.Name,stu.Age)
	return str
}
func main() {
	var stu = Student{
		Name: "张三",
		Age:  20,
	}
	fmt.Println(stu)  //表示直接输出结构体 {张三 20}
	fmt.Println(&stu) //按照定义的格式输出 [姓名=张三,年龄=20]
}
