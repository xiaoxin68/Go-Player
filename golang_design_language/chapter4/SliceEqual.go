package main

import (
	"bytes"
	"fmt"
)

func main() {
	//和数组不用的是，slice无法做比较！！！
	//但是bytes提供了Equal方法比较[]byte型的切片，其他类型切片比较得自己写
	a := []byte{'a','b'}
	b := []byte{'a','b'}
	e := bytes.Equal(a,b)
	fmt.Println(e)

	x := []string{"sfg","fjdfa"}
	y := []string{"sfg","fjdfa"}
	z := eql(x,y)
	fmt.Println(z)
}

func eql(x,y []string)bool  {
	if len(x) != len(y){
		return false
	}
	for i:= range x{
		if x[i] != y[i]{
			return false
		}
	}
	return true
}
