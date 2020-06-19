package main

import "fmt"

//Copy：可对其切片进行深拷贝
func main() {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{7, 8, 9}
	i := copy(s1, s2)
	fmt.Println(i)
	fmt.Println(s1)
	fmt.Println(s2)

	j := copy(s2, s1[2:])
	fmt.Println(j)
	fmt.Println(s1)
	fmt.Println(s2)

	k := copy(s1[:2], s2[1:])
	fmt.Println(k)
	fmt.Println(s1)
	fmt.Println(s2)
}
