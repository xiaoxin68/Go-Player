package main

import "fmt"

func main() {
	ch := make(chan bool) //创建通道
	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("goroutine:", i)
		}
		ch <- true //写入
	}()
	data := <-ch //读取：在ch中没有数据时，这个操作会被阻塞
	fmt.Println(data)
	fmt.Println("main end...")
}
