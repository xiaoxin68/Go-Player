package main

import (
	"fmt"
)

func main() {
	//go语言允许在if语句中直接定义变量
	//即时只有一条语句，也要带上{}
	if age := 20; age > 18 {
		fmt.Printf("年龄大于18了")
	}
}
