package main

import "fmt"

//如果结构体的字段类型是：指针，slice，map的零值都是nil，即还没有分配空间
//如果需要使用这样的字段，需要先make才能使用
type Person struct {
	Name string
	Age int
	Scores [5]float64
	ptr *int //指针
	slice []int //切片
	map1 map[string]string //切片
 }
func main() {
	var person Person
	
	fmt.Println(person)
	
	person.Name="张三"
	person.Age=18
	//数组
	var myScore =  [5]float64{89.90,90.89,90.89,90.89,67.89}
	person.Scores = myScore
	//指针
	i:=10
	person.ptr = &i
	//切片
	person.slice = make([]int,10)
	person.slice[0] = 100
	//map
	person.map1 = make(map[string]string,5)
	person.map1["zhangsan"] = "haha"
	
	fmt.Println(person)
	
	modifyStruct(person)
	
	fmt.Println(person)
}

//结构体是值类型
func modifyStruct(person Person)  {
	person.Name = "李四"
	person.Age=20
	//修改数组
	person.Scores[0] = 100.00
	//修改指针
	var pt int = 100
	person.ptr = &pt
	//修改切片
	person.slice[0] = 200
	person.slice[1] = 200
	//修改map
	person.map1["zhangsan"] = "heihei"
	person.map1["lisi"] = "loulou"
}