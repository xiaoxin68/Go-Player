package main

import (
	"fmt"
	"strings"
)

func main() {
	f := makeSuffix(".jpg")
	fmt.Println(f("hello.jpg"))
	fmt.Println(f("zxx"))
	
	f1 := makeSuffix(".png")
	fmt.Println(f1("hello.png"))
	fmt.Println(f1("zxx"))
}

//使用闭包的好处：传入一次后缀就可以反复使用
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			name += ".jpg"
		}
		return name
	}
}
