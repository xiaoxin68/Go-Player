package main

import "fmt"

type Student struct {
	name string
	age int
}
//这里如果我们不使用 new 函数为自定义类型分配空间（将第 12 行注释），就会报错：
func main() {
	var s *Student
	s = new(Student) //分配空间
	s.name ="dequan"
	fmt.Println(s)
}
//new 函数，它返回的永远是类型的指针，指针指向分配类型的内存地址

/*Go语言中的 new 和 make 主要区别如下：
make 只能用来分配及初始化类型为 slice、map、chan 的数据。new 可以分配任意类型的数据；
new 分配返回的是指针，即类型 *Type。make 返回引用，即 Type；
new 分配的空间被清零。make 分配空间后，会进行初始化；*/