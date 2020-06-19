package main

import (
	"fmt"
	"runtime"
)

//go1.8后，默认让程序运行在多个核上，可以不用设置了
//go1.8前，还是要设置一下
func main() {
	//获取当前系统cpu数量
	num := runtime.NumCPU() //NumCPU返回本地机器的逻辑CPU个数。
	//这里设置num-1的cpu运行go程序
	runtime.GOMAXPROCS(num-1)  //GOMAXPROCS设置可同时执行的最大CPU数
	fmt.Println("num=",num)
}
