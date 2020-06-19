package main

import "fmt"

func main() {
	fmt.Println("sum=", sumFun(1, 2))//sum= 3
}

func sumFun(n1, n2 float32) float32 {
	fmt.Printf("n1 type = %T\n", n1)//n1 type = float32
	return n1 + n2
}
