package main

import (
	"fmt"
	"time"
)

func main() {
	begin := time.Now().Unix()
	for i:=1 ;i<=10000;i++{
		fmt.Println(i)
	}
	end := time.Now().Unix()
	fmt.Println("程序执行耗费秒数：",end-begin)
}
