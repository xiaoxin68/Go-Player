package main

import (
	"fmt"
	"unicode"
)

func main() {
	var ch = 'a'
	var ch2 = '1'
	//判断是否为字母：
	fmt.Println(unicode.IsLetter(ch))
	//判断是否为数字：
	fmt.Println(unicode.IsDigit(ch2))
	//判断是否为空白符号：
	fmt.Println(unicode.IsSpace(ch))
}
