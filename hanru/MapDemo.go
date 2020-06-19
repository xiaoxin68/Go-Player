package main

import "fmt"

func main() {
	//m1 := make(map[string]string,2)
	m1 := make(map[string]string)
	m1["张三"] = "护士"
	m2 := m1
	fmt.Println("=======修改前========")
	fmt.Println(m1, "len=", len(m1))
	fmt.Println(m2, "len=", len(m2))

	fmt.Println("=======修改内容========")
	m1["张三"] = "医生"
	fmt.Println(m1, "len=", len(m1))
	fmt.Println(m2, "len=", len(m2))

	fmt.Println("=======增减元素，未超过size========")
	m1["李四"] = "飞行员"
	fmt.Println(m1, "len=", len(m1))
	fmt.Println(m2, "len=", len(m2))

	fmt.Println("=======增减元素，超过size========")
	m1["王二"] = "厨师"
	fmt.Println(m1, "len=", len(m1))
	fmt.Println(m2, "len=", len(m2))
}
