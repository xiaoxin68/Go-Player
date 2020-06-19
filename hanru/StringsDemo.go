package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "hello world"
	fmt.Println(str)
	//join
	arr := []string{"sdas", "asdsa", "Asff"}
	strings.Join(arr, "-")
	fmt.Println(arr)
	//split
	str2 := "hello,sgf,asd,dafsd,dfz"
	arr2 := strings.Split(str2, ",")
	fmt.Println(arr2)
	//substring
	str3 := str[:5] //利用切片操作的方式
	fmt.Println(str3)
}