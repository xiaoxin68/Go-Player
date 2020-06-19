package main

import (
	"fmt"
	"time"
)

//注意：以下程序存在问题
//1、编写一个函数，来计算各个数的阶乘，并放入到map中
//2、我们启动的协诚多个，统计的价结果将放入到map中
//3、map应定义成全局的

var (
	myMap = make(map[int]int, 10) //fatal error: concurrent map writes
)

func test1(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	myMap[n] = res
}
func main() {
	//初始化map
	for i := 1; i <= 200; i++ {
		go test1(i)
	}
	//睡眠10s
	time.Sleep(time.Second * 10)
	//输出结果
	for idx, val := range myMap {
		fmt.Printf("map[%d]=%d\n", idx, val)
	}
}
//Found 2 data race(s)