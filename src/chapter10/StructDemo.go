package main

import "fmt"

//结构体的所有字段在内存中是连续的
type Point struct {
	x int
	y int
}
type Rect struct {
	leftUp,rightDown Point
}
type Rect2 struct {
	leftUp,rightDown *Point
}
func main() {
	r1 := Rect{Point{1,2},Point{3,4}}
	//r1有四个int，在内存中连续分布
	fmt.Printf("r1.leftUp.x 地址=%p r1.leftUp.y 地址=%p r1.rightDown.x 地址=%p r1.rightDown.y 地址=%p\n",
		&r1.leftUp.x,&r1.leftUp.y,&r1.rightDown.x,&r1.rightDown.y)
	
	//r2有两个*Point类型，这两个*Point类型的地址也是连续的
	//但是它们指向的地址不一定连续
	r2 := Rect2{&Point{10,20},&Point{30,40}}
	fmt.Printf("r2.leftUp本身的地址=%p  r2.rightDown本身的地址=%p\n",
		&r2.leftUp,&r2.rightDown)
	
	fmt.Printf("r2.leftUp指向的地址=%p  r2.rightDown指向的地址=%p\n",
		r2.leftUp,r2.rightDown)
}
