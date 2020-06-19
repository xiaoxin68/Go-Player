package main

import "fmt"

func main() {
	a := sum1
	b := sum2
	fmt.Println(myFun1(a, 1, 2)) //3
	fmt.Println(myFun2(b,1,2,3,4,5,6)) //21
}

func sum1(n1 int, n2 int) int {
	return n1 + n2
}
type mySum1 func(int, int) int  //别名
func myFun1(funVar mySum1, num1 int, num2 int) int {
	return funVar(num1, num2)
}

func sum2(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}
type mySum2 func(int, ...int) int
func myFun2(funVar mySum2, num1 int, args ...int) int {
	return funVar(num1, args...)
}