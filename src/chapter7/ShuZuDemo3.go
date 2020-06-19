package main

import "fmt"

//数组传递是值传递
//arr= [1 2 3]
//arr= [1 2 3]
//arr= [1 88 3]
func main() {
	var arr = [...]int{1, 2, 3}
	fmt.Println("arr=", arr)
	test(arr)
	fmt.Println("arr=", arr)
	test2(&arr)
	fmt.Println("arr=", arr)
}

func test(arr [3]int) {
	arr[1] = 88
}

func test2(arr *[3]int) {
	(*arr)[1] = 88
}
