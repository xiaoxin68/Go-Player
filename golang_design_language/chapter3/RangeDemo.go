package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "hello北京"
	n := 0
	for range s {
		n++
	}
	fmt.Println(n) //7

	r := []rune(s)
	fmt.Printf("%x\n", r) //[68 65 6c 6c 6f 5317 4eac]

	fmt.Println(string(r)) //hello北京

	fmt.Println(string(0x4eac)) //京

	fmt.Println(string(65)) //A  而不是 65

	fmt.Println(strconv.FormatInt(65,10))  //65



	fmt.Println(string(1234567)) //�  非法
}
