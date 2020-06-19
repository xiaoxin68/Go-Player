package main

import "fmt"

func main() {
	ch := make(chan int)
	go sendDatas(ch)
	for v := range ch{
		fmt.Println("读取数据：",v)
	}
	fmt.Println("main end...")
}
func sendDatas(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}