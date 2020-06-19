package main

import "fmt"

func main() {
	ch := make(chan int)
	go sendData(ch)
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("已经读取了所有的数据")
			break
		}
		fmt.Println("读取到的数据为：", v)
	}
}

func sendData(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}
