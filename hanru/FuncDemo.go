package main

import "fmt"

//函数可以作为参数
//匿名函数
func main() {
	//函数可以作为参数
	res1 := f1(10, 20, 30, add)
	fmt.Println(res1) //10+(20+30)
	//函数可以作为参数
	res2 := f1(10, 20, 30, sub)
	fmt.Println(res2) //10+(20-30)

	//匿名函数
	res3 := func(a, b int) int {
		return a - b
	}
	fmt.Println(res3(10, 20)) //10-20
	//匿名函数
	res4 := func(s1, s2 string) string {
		return s1 + s2
	}("hello", "world")
	fmt.Println(res4) //hello+world
}

func add(m, n int) int {
	return m + n
}
func sub(m, n int) int {
	return m - n
}

//回调函数
func f1(a, b, c int, fun func(int, int) int) int {
	res := fun(b, c)
	return res + a
}
