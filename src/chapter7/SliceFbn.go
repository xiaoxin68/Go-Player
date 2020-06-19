package main

import "fmt"

func main() {
	uint64s := fbn(10)
	fmt.Println(uint64s)  //[1 1 2 3 5 8 13 21 34 55]
	uint64s2 := fbn2(10)
	fmt.Println(uint64s2)
}

//动态规划写法
func fbn(n int) []uint64 {
	var res []uint64 = make([]uint64, n)
	res[0] = 1
	res[1] = 1
	var num1, num2 uint64 = 1, 1
	for i := 2; i < n; i++ {
		res[i] = num1 + num2
		num1 = num2
		num2 = res[i]
	}
	return res
}

//递归写法
func fbn2(n int) []uint64 {
	var res []uint64 = make([]uint64, n)
	res[0] = 1
	res[1] = 1
	for i := 2; i < n; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return res
}
