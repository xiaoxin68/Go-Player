package main

import "fmt"

func main() {
	var list1 = []string{"ds2f", "fds", "SAD"}
	var list2 = []string{"dsf", "asd", "Afa"}
	var list3 = []string{"aSD", "fds", "SAD"}
	//如果现在希望list1、list2、list3作为key
	//1、把key映射到字符串
	add(list1)
	add(list2)
	add(list2)
	add(list3)
	fmt.Println(count(list1)) //1
	fmt.Println(count(list2)) //2
	fmt.Println(m)            //map[["aSD" "fds" "SAD"]:1 ["ds2f" "fds" "SAD"]:1 ["dsf" "asd" "Afa"]:2]
}

var m = make(map[string]int)

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}
func add(list []string) {
	m[k(list)]++
}
func count(list []string) int {
	return m[k(list)]
}
