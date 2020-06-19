package main

import "fmt"

//键盘的输入
func main() {
	var name string
	var age byte
	var sal float32
	var isPass bool

	//方式一：Scanln
	//fmt.Println("请输入姓名：")
	//fmt.Scanln(&name) //注意这里要用地址，这样对于它的修改才会修改到name
	//
	//fmt.Println("请输入年龄：")
	//fmt.Scanln(&age)
	//
	//fmt.Println("请输入薪水：")
	//fmt.Scanln(&sal)
	//
	//fmt.Println("请输入是否通过：")
	//fmt.Scanln(&isPass)
	//
	////name= 啦啦啦 ,age= 23 ,sal= 123.43 ,isPass= true
	//fmt.Println("name=", name, ",age=", age, ",sal=", sal, ",isPass=", isPass)

	//方式二：Sscanf
	fmt.Println("你的姓名，年龄，薪水，是否通过考试分别为：【用空格隔开】")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	//姓名:tom         年龄:23         薪水:123.34     是否通过考试:false
	fmt.Printf("姓名:%v \t 年龄:%v \t 薪水:%v \t 是否通过考试:%v \n",name,age,sal,isPass)
}
