package main

import "fmt"

//注意：下面写法是错误的，需要先make再使用,或者创建的时候就初始化
//如果不想make，也可以采用append
func main() {
	/*var a []int
	a[0] = 10
	fmt.Println(a)*/

	var len = 10
	var a []int = make([]int, len)
	fmt.Println(a == nil)
	a[0] = 10
	fmt.Println(a)

	var b []int
	b = append(b, 20)
	fmt.Println(b)
}
