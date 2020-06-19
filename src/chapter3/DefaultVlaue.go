package main

import "fmt"

func main() {
	var i int
	var j float32
	var k bool
	var m byte
	var n string
	//%v表示按照变量的值输出
	// //%d %c %s %T %f分别是以 整型 字符 字符串 类型 浮点型 输出；
	fmt.Printf("i=%d,j=%f,k=%v,m=%d,n=%v", i, j, k, m, n) //i=0,j=0.000000,k=false,m=0,n=
	
}
