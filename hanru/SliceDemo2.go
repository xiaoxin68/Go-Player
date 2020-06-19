package main

import "fmt"

func main() {
	s1 := make([]int, 1, 2)
	s1[0] = 1
	s2 := s1
	fmt.Println("=======修改前========")
	fmt.Println(s1) //[1]
	fmt.Println(s2) //[1]

	fmt.Println("=======修改内容========")
	s1[0] = 100
	fmt.Println(s1) //[100]
	fmt.Println(s2) //[100]

	fmt.Println("=======append未扩容下修改内容========")
	s1 = append(s1, 10)
	s1[0] = 200
	fmt.Println(s1) //[200 10]
	fmt.Println(s2) //[200]

	fmt.Println("=======append扩容下修改内容========")
	s1 = append(s1, 11, 12)
	s1[0] = 300
	fmt.Println(s1) //[300 10 11 12]
	fmt.Println(s2) //[200]
}
