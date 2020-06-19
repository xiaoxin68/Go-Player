package main

import "fmt"

func main() {
	var n int = 6
	for i := 0; i <= n; i++ {
		fmt.Printf("%d\t+\t%d\t=\t%d\n", i, n-i, n)
	}
	
	//注意go语言里面没有while和do while语法
}
