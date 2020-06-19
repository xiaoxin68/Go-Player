package main

import "fmt"

func main() {
	//进制转化
	var i int = 13
	fmt.Printf("二进制表示：%b \n", i)
	fmt.Printf("十进制表示：%d \n", i)
	fmt.Printf("八进制表示：%o \n", i)
	fmt.Printf("十六进制【小写字母】表示：%x \n", i)
	fmt.Printf("十六进制【大写字母】表示：%X \n", i)
	
	//位运算
	a := 1 >> 2
	b := -1 >> 2
	c := 1 << 2
	d := -1 << 2
	fmt.Printf("a=%v,b=%v,c=%v,d=%v\n", a, b, c, d) //a=0,b=-1,c=4,d=-4
	
	
}
