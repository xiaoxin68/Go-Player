package main

import "fmt"

func main() {
	//1、方式一：先声明，再make，再初始化
	var a map[string]string
	a = make(map[string]string,10)
	a["idx1"] = "张三"
	fmt.Println(a)
	
	//2、方式二：声明就直接make，再初始化
	b := make(map[string]string,10)
	b["idx1"] = "北京"
	fmt.Println(b)
	
	//3、方式三：声明就直接初始化
	c := map[string]string{
		"idx1":"星期一",
		"idx2":"星期二",
		"idx3":"星期三",
	}
	fmt.Println(c)
}
