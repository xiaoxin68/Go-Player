package main

import "fmt"

func main() {
	stack := []string{"af", "Af", "sad", "ASf", "SAfa"}
	//push
	stack = append(stack, "weq")
	//top
	top := stack[len(stack)-1]
	fmt.Println(top) //weq
	//pop
	pop := stack[:len(stack)-1]
	fmt.Println(pop) //[af Af sad ASf SAfa]

	fmt.Println(stack)      //[af Af sad ASf SAfa weq]
	res := remove(stack, 3) //[af Af sad SAfa weq]
	fmt.Println(res)
}
func remove(strs []string, i int) []string {
	//移除第i个元素
	copy(strs[i:], strs[i+1:])
	return strs[:len(strs)-1]
}
