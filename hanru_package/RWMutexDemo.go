package main

import (
	"fmt"
	"sync"
	"time"
)

//读写锁
var rwMutex sync.RWMutex
var wg2 sync.WaitGroup

func main() {
	wg2.Add(4)
	go readData(1)
	go readData(2)
	go writeData(3)
	go writeData(4)
	wg2.Wait()
	fmt.Println("main函数结束~~~~~~~~~~~·")
}

//读操作可以同时进行
func readData(i int) {
	defer wg2.Done()
	rwMutex.RLock() //读操作上锁
	fmt.Println(i, "开始读：read start")
	fmt.Println(i, "正在读取数据：reading。。。")
	time.Sleep(time.Second)
	fmt.Println(i, "结束读：read end")
	rwMutex.RUnlock() //读操作解锁
}

//写操作不可以同时进行
func writeData(i int) {
	defer wg2.Done()
	rwMutex.Lock() //写操作上锁
	fmt.Println(i, "开始写：write start")
	fmt.Println(i, "正在写数据：writeing。。。")
	time.Sleep(time.Second)
	fmt.Println(i, "结束写：write end")
	rwMutex.Unlock() //写操作解锁

}
