package main

import "fmt"

func main() {
	var a interface{}
	var b float64 = 12.34
	a = b
	if y, ok := a.(float64); ok {
		fmt.Println("convert success")
		fmt.Printf("y的类型是%T，y=%v", y, y)
	} else {
		fmt.Println("convert fail")
	}
	fmt.Println("继续执行")
}
