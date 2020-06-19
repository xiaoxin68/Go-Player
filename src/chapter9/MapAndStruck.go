package main

import "fmt"

type Stu struct {
	name string
	age int
	addr string
}
func main() {
	students := make(map[string]Stu,5)
	students["no1"] = Stu{"Tom",18,"上海"}
	students["no2"] = Stu{"Lili",19,"北京"}
	fmt.Println(students)
	
	for key,val := range students {
		fmt.Println("学号：",key)
		fmt.Printf("\t姓名：%v\n",val.name)
		fmt.Printf("\t年龄：%v\n",val.age)
		fmt.Printf("\t住址：%v\n",val.addr)
	}
}
