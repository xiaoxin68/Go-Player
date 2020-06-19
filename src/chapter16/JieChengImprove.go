package main

import (
	"fmt"
	"sync"
	"time"
)

//不同goroutine之间如何通讯
//1）全局变量的互斥锁【不太好】
//2）使用管道channel来解决


//全局变量的互斥锁【不太好】
var (
	myMap2 = make(map[int]int, 10)
	
	//声明一个互斥的全局锁
	//lock是一个全局的互斥锁
	//sync是包：synchronize同步
	//Mutex：是互斥
	lock sync.Mutex
)

func test2(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	
	//加锁
	lock.Lock()
	myMap2[n] = res
	//解锁
	lock.Unlock()
}
func main() {
	//初始化map
	for i := 1; i <= 20; i++ {
		go test2(i)
	}
	
	//这个时间的设置不好估算。时间设置的短了，主线程执行完了协程还没执行完；时间长了等待时间太长了不好
	time.Sleep(time.Second * 10)
	
	//输出结果
	//这里也需要加锁：因为我们程序从设计上可以知道10秒就执行完所有协诚，但是主线程并不知道
	//因此底层仍然可能出现资源争夺，因此加入互斥锁即可解决问题
	lock.Lock()
	for idx, val := range myMap2 {
		fmt.Printf("map2[%d]=%d\n", idx, val)
	}
	lock.Unlock()
}
