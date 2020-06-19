package main

import "fmt"

type People struct {
	Name string
	Age  int
}

func main() {
	//方式一
	p1 := People{"小白", 19}
	fmt.Println(p1)
	
	//方式二
	var p2 People
	p2.Name = "小花"
	p2.Age = 2
	fmt.Println(p2)
	
	//方式三
	var p3 *People = new(People)
	(*p3).Name = "小黄" //也可以这么写：p3.Name = "小黄"；这是go的设计者为了简化，允许这么写
	(*p3).Age = 11    //也可以这么写：p3.Age = 11
	fmt.Println(*p3)
	
	//方式四
	var p4 *People = &People{}
	(*p4).Name = "小粉" //也可以这么写：p4.Name = "小粉"
	(*p4).Age = 13    //也可以这么写：p4.Age = 13
	fmt.Println(*p4)
	
	var p5 *People = &p1 //这个时候改变p5，相应的p1也会改变
	(*p5).Name = "小小"
	fmt.Println(p1)
	fmt.Println(*p5)
}
