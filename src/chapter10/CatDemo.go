package main

import "fmt"

type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}

func main() {
	cats := make(map[string]Cat, 10)
	cats["小白"] = Cat{"小白", 10, "白色", "睡觉"}
	cats["小花"] = Cat{"小花", 12, "花色", "爬树"}
	
	var name string
	fmt.Println("请输入小猫的名字：")
	fmt.Scanln(&name)
	val, ok := cats[name]
	if ok {
		fmt.Println("姓名：", val.Name)
		fmt.Println("年龄：", val.Age)
		fmt.Println("花色：", val.Color)
		fmt.Println("爱好：", val.Hobby)
	} else {
		fmt.Println("猫猫不存在")
	}
}
