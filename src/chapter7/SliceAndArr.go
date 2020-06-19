package main

import "fmt"

func main() {
	var slice []int
	var arr = [5]int{1, 2, 3, 4, 5}
	slice = arr[:]
	slice2 := slice
	slice2[0] = 10
	fmt.Println("arr=", arr)       //arr= [10 2 3 4 5]
	fmt.Println("slice=", slice)   //slice= [10 2 3 4 5]
	fmt.Println("slice2=", slice2) //slice2= [10 2 3 4 5]
	
	test01(slice)
	
	fmt.Println("arr=", arr)       //arr= [20 2 3 4 5]
	fmt.Println("slice=", slice)   //slice= [20 2 3 4 5]
	fmt.Println("slice2=", slice2) //slice2= [20 2 3 4 5]
	
	test02(arr)
	
	fmt.Println("arr=", arr)       //arr= [20 2 3 4 5]
	fmt.Println("slice=", slice)   //slice= [20 2 3 4 5]
	fmt.Println("slice2=", slice2) //slice2= [20 2 3 4 5]
}
func test01(slice []int) {
	slice[0] = 20
}

func test02(arr [5]int) {
	arr[0] = 30
}
