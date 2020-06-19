package main

import "fmt"

func main()  {
	fmt.Println(Fib(10))
	fmt.Println(Fib2(10))
}
//非递归写法
func Fib(n int)  int{
	if n <= 0 {
		return 0
	}else if n == 1 {
		return 1
	}else {
		n1 := 0
		n2 := 1
		var res int
		for i:=2;i<=n ;i++  {
			res = n1 + n2
			n1 = n2
			n2 = res
		}
		return res
	}
}
//递归写法
func Fib2(n int)  int{
	if n <= 0 {
		return 0
	}else if n == 1 {
		return 1
	}else {
		return Fib2(n-1) + Fib2(n-2)
	}
}