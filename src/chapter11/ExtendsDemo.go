package main

import "fmt"

type A struct {
	Name string
	age  int
}

func (a *A) sayHi() {
	fmt.Println("A sayHi", a.Name)
}
func (a *A) SayHello() {
	fmt.Println("A SayHello", a.Name)
}

type B struct {
	Name string //如果子类包含父类同名的属性
	A
}

func (b *B) SayHello() {
	fmt.Println("B SayHello", b.Name)
}
func main() {
	var b B
	b.A.Name = "name1" //可调用父类公有属性
	b.A.age = 10       //可调用父类私有属性
	b.A.SayHello()     //可调用父类公有方法
	b.A.sayHi()        //可调用父类私有方法
	
	//简化
	//var b1 B
	//b1.Name = "name2"
	//b1.age = 20
	//b1.SayHello()
	//b1.sayHi()
	
	var b2 B
	b2.Name = "name3" //赋值的是B里面的Name；而A中的Name=""
	b2.A.Name = "name4"
	b2.age = 30       //赋值的是A里面的age
	b2.sayHi()        //调用的A里面的方法  :   A sayHi name4
	b2.SayHello()     //调用的B里面的方法  :  B SayHello name3
}
