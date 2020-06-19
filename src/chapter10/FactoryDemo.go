package main

import (
	"./model"
	"fmt"
)

func main() {
	var stu = model.Student{"张三", 19}
	fmt.Println(stu)
	fmt.Println(stu.Name, stu.Age)
	
	//model包中的结构体首字母大写
	var stu2 = model.NewStudent("tom", 21)
	fmt.Println(*stu2)
	fmt.Println(stu2.Name, stu2.Age)
	
	//model包中的结构体中某字段首字母小写
	var stu3 = model.NewStudent1("lili", 10)
	fmt.Println(*stu3)
	fmt.Println(stu3.Name, stu3.GetAge())
}
