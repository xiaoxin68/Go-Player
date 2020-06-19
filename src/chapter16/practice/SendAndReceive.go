package main

import "fmt"

//chan<- int:这样ch只能进行写操作
func send(ch chan<- int, exitChan chan struct{}) {
	for i := 1; i <= 50; i++ {
		ch <- i
	}
	close(ch)
	var a struct{}
	exitChan <- a
}

//ch <-chan int:这样ch只能进行读操作
func recv(ch <-chan int, exitChan chan struct{}) {
	for {
		val, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(val)
	}
	var a struct{}
	exitChan <- a
}
func main() {
	var ch chan int
	ch = make(chan int, 10)
	exitChan := make(chan struct{}, 2)

	go send(ch, exitChan)
	go recv(ch, exitChan)

	var total = 0
	for _ = range exitChan {
		total++
		if total == 2 {
			break
		}
	}
	fmt.Println("结束")
}
