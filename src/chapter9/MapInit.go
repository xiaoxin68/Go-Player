package main

import "fmt"

//map是无序的
func main() {
	var a map[string]string
	fmt.Println(a) //map[]
	
	//map在使用前需要先make一下，分配空间
	a = make(map[string]string, 10)
	
	a["idx1"] = "hello"
	a["idx2"] = "world"
	fmt.Println(a) //map[idx1:hello idx2:world]
	a["idx1"] = "hello!!!"
	fmt.Println(a) //map[idx1:hello!!! idx2:world]
}
