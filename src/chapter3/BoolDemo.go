package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i bool = true
	fmt.Println("i=", i)
	fmt.Println("i占用的字节是:", unsafe.Sizeof(i)) //i占用的字节是: 1
	
	//不可用0或非0来表示false或true
	var j bool
	fmt.Println("j=", j) //默认为false
}
