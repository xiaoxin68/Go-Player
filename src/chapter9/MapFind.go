package main

import "fmt"

func main() {
	var a = make(map[string]string, 10)
	a["idx1"] = "hello"
	a["idx2"] = "world"
	a["idx3"] = "zxx"
	
	val,ok:=a["idx2"]
	if ok {
		fmt.Println(val)
	}else {
		fmt.Println("map中不存在该元素")
	}
}
