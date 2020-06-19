package main

import "fmt"

//只输出ok2
func main() {
	var num int = 10
	if num < 9 && test() { //第一个条件为false，第二个条件就不会判断
		fmt.Println("ok1")
	}

	if num > 9 || test() { //第一个条件为true，第二个条件就不会判断
		fmt.Println("ok2")
	}
}
func test() bool {
	fmt.Println("###########")
	return true
}
