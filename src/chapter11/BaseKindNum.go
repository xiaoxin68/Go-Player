package main

import "fmt"

type Monster struct {
	Name string
	Age  int
}
type E struct {
	Monster
	int
}
type F struct {
	Monster
	int
	i int
}

func main() {
	var e = E{
		Monster: Monster{
			Name: "张三",
			Age:  15,
		},
		int: 25,
	}
	fmt.Println(e)
	
	var f F
	f.Name = "李四"
	f.Age = 10
	f.int = 20 //基本类型可以直接这么赋值
	f.i = 30
	fmt.Println(f)
}
