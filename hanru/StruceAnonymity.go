package main

import "fmt"

//匿名结构体和匿名字段
func main() {
	//匿名结构体
	anonymityStruct := struct {
		name string
		age  int
	}{
		name: "张三",
		age:  13,
	}
	fmt.Println(anonymityStruct, anonymityStruct.name, anonymityStruct.age)

	annoy := anonymityFiled{"啦啦啦", 100}
	fmt.Println(annoy, annoy.string, annoy.int)
}

type anonymityFiled struct {
	string
	int
}
