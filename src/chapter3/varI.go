package main

import "fmt"

//定义全局变量
var l1 = 100
var l2 = 12.3
var l3 = "zhangsan"
var(
	r1 = 200
	r2 = 13.3
	r3 = "zhangsan"
)

func main() {
	var i int //定义变量
	i = 10 //给变量赋值
	fmt.Println("i=",i) //使用变量
	
	//新建一个变量，如果不赋值，使用默认值
	var j int
	fmt.Println("j=",j) //int类型如果不初始化，默认值为0
	
	//根据值自行判断类型
	var k = 1.01
	fmt.Println("k=",k)
	
	//等价于
	// var name string
	//name = "tom"
	name := "tom"
	fmt.Println("name=",name)
	
	//n1,n2,n3 := 10,10.23,"zhnag"
	//或者
	var n1,n2,n3=10,10.23,"zhnag"
	fmt.Println("n1=",n1,"\t n2=",n2,"\t n3=",n3)
	
	fmt.Println("l1=",l1,"\t l2=",l2,"\t l3=",l3)
	fmt.Println("r1=",r1,"\t r2=",r2,"\t r3=",r3)
}
