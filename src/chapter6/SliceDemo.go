package main

import "fmt"

func main() {
	res := sliceFun(10, 11, 23, 23)
	fmt.Println(res) //67
}

//go支持可变参数
func sliceFun(num int, args ...int) int {
	sum := num
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}
