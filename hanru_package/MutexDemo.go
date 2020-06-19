package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//临界资源
var ticket = 10
var mutex sync.Mutex     //互斥锁
var waitG sync.WaitGroup //等待一组协程的结束
func main() {
	waitG.Add(4) //等待的协程数
	go saleTicket("窗口1")
	go saleTicket("窗口2")
	go saleTicket("窗口3")
	go saleTicket("窗口4")
	waitG.Wait() //表示main协程进入等待，意味着阻塞
}
func saleTicket(name string) {
	defer waitG.Done() //等待协程数减一
	rand.Seed(time.Now().UnixNano())
	for {
		mutex.Lock() //加锁
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出：", ticket)
			ticket--
		} else {
			mutex.Unlock() //条件不满足也要解锁
			fmt.Println(name, "售罄，没票了。。。")
			break
		}
		mutex.Unlock() //解锁
	}
}
