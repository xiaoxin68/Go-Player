package main

import "fmt"

type Point struct {
	x int
	y int
}
func main() {
	var a interface{}
	var point Point = Point{10,20}
	a = point
	var b Point
	b = a.(Point)  //类型断言：如果类型转换失败的话，会报panic错误
	fmt.Println(b)
}
