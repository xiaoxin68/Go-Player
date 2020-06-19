package main

import "fmt"

func main()  {
	var i int
	i = 20
	//i = 10.1 //.\vaI2.go:8:4: constant 10.1 truncated to integer  不能改变数据类型
	//i := 30  变量在同一作用域不能重名
	fmt.Println(i)

}
