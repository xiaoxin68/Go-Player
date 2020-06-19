package main

import "fmt"

func main() {
	var a = make(map[string]string, 10)
	//新增
	a["idx1"] = "hello"
	a["idx2"] = "world"
	a["idx3"] = "zxx"
	fmt.Println(a)
	//更新
	a["idx1"] = "hello!!!"
	fmt.Println(a)
	//删除
	delete(a,"idx2")
	fmt.Println(a)
	//删除map中所有的元素：要么遍历，一个个删除；要么重新make一下
	a = make(map[string]string)
	fmt.Println(a)
}
