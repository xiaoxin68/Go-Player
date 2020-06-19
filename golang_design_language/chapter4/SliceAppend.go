package main

import "fmt"

func main() {
	a := []int{1,2,3}
	a = append(a,4)  //内置的appen函数
	a = appendInt(a,5) //自己实现的append函数
	fmt.Println(a) //[1 2 3 4 5]

	//演示用量不够时的扩容
	var x,y []int
	for i:=0;i<10;i++ {
		y = appendInt(x,i)
		fmt.Printf("%d cap=%d  %v\n",i,cap(y),y)
		x = y
	}
}
func appendInt(x []int,y int)[]int  {
	var z []int
	zlen := len(x)+1
	if zlen <= cap(x) {
		z = x[:zlen]
	}else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int,zlen,zcap)
		copy(z,x)
	}
	z[len(x)] = y
	return z
}

/*0 cap=1  [0]
1 cap=2  [0 1]
2 cap=4  [0 1 2]
3 cap=4  [0 1 2 3]
4 cap=8  [0 1 2 3 4]
5 cap=8  [0 1 2 3 4 5]
6 cap=8  [0 1 2 3 4 5 6]
7 cap=8  [0 1 2 3 4 5 6 7]
8 cap=16  [0 1 2 3 4 5 6 7 8]
9 cap=16  [0 1 2 3 4 5 6 7 8 9]*/