package main

import "fmt"

type address struct {
	host string
	port int
}
func main() {
	//可比较的结构体类型可以作为map的键
	hits := make(map[address]int)
	hits[address{"sada",12}]=20

	hits2 := make(map[address]int)
	hits2[address{"sada",12}]=20

	fmt.Println(address{"sada",12} == address{"sada",12})//true

	//但是slice不可以作为map的键，因为slice不可比较
	//Invalid operation: []int{1,2,3} == []int{1,2,3} (operator == is not defined on []int)
	//fmt.Println([]int{1,2,3} == []int{1,2,3})
}
