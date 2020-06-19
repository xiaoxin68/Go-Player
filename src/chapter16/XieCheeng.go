package main

import (
	"fmt"
	"strconv"
	"time"
)

//编写一个函数，每隔一秒删除hello world
func test()  {
	for i:=1;i<=100 ; i++ {
		fmt.Println("test() hello world"+strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go test() //开启一个协程
	
	//注意：如果主线程退出了，即时协程还没有执行完，也会退出！！！！
	for i:=1;i<=10 ; i++ {
		fmt.Println("main() hello world"+strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
