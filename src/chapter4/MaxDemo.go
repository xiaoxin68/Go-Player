package main

import "fmt"

//go不支持三元运算符
func main() {
	var i, j int = 12, 23
	var max int
	if i >= j {
		max = i
	} else {
		max = j
	}
	fmt.Println("max:", max)
}
