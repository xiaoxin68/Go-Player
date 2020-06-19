package main

import "fmt"

func main()  {
	
	//用for循环模拟while
	i := 1
	for{
		if i > 5 {
			break
		}
		fmt.Println("hello")
		i++
	}
	
	//用for循环模拟do while
	j := 1
	for{
		fmt.Println("world")
		j++
		if j > 5 {
			break
		}
	}
}
