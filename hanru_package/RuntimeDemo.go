package main

import (
	"fmt"
	"runtime"
	"time"
)

func init() {
	//设置逻辑CPU的数量
	n := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(n)
}
func main() {
	//GOROOT
	fmt.Println(runtime.GOROOT())
	//操作系统
	fmt.Println(runtime.GOOS)
	//逻辑CPU数量
	fmt.Println(runtime.NumCPU())

	/*go func() {
		for i:=0;i<10;i++ {
			fmt.Println("gotoutine...")
		}
	}()

	for i:=0;i<10;i++ {
		//礼让：使当前go程放弃处理器，以让其它go程运行。它不会挂起当前go程，因此当前go程未来会恢复执行。
		runtime.Gosched()
		fmt.Println("main...")
	}*/

	go func() {
		fmt.Println("go协程开始")
		fun()
		fmt.Println("go协程结束")
	}()

	time.Sleep(time.Second * 3)
}

func fun() {
	defer fmt.Println("defer...")
	//return   //终止当前函数
	runtime.Goexit() //终止当前goroutine
	fmt.Println("fun函数")
}
