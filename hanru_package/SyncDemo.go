package main

import (
	"fmt"
	"sync"
)

//WaitGroup用于等待一组线程的结束。父线程调用Add方法来设定应等待的线程的数量。
//每个被等待的线程在结束时应调用Done方法。同时，主线程里可以调用Wait方法阻塞至所有线程结束。
var wg sync.WaitGroup

func main() {
	//调用Add方法来设定应等待的线程的数量
	wg.Add(2)
	go fun1()
	go fun2()

	fmt.Println("main进入阻塞状态，等待wg中的子协程结束")
	wg.Wait()  //表示main协程进入等待，意味着阻塞
	fmt.Println("main解除阻塞")
}

func fun1() {
	for i := 1; i <= 10; i++ {
		fmt.Println("fun1()函数打印。。。", i)
	}
	wg.Done() //底层是wg.Add(-1)
}

func fun2() {
	defer wg.Add(-1)
	for i := 1; i <= 10; i++ {
		fmt.Println("fun2()函数打印==========", i)
	}
}
