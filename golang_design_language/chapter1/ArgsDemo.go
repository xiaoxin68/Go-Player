package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ { //var Args []string是一个切片
		s += sep + os.Args[i]
		sep = ""
	}
	fmt.Println(s)

	//func Join(a []string, sep string) string
	//将一系列字符串连接为一个字符串，之间用sep来分隔
	fmt.Println(strings.Join(os.Args[1:],","))

	fmt.Println(os.Args[0])
}
