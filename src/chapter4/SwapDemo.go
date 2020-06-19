package main

import "fmt"

func main() {
	var a, b int = 10, 20
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("a=", a, ",b=", b) //a= 20 ,b= 10
}
