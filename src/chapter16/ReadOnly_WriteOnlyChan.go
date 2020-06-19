package main

import "fmt"

func main() {
	//只写的chan
	var readChan  chan<- int
	readChan = make(chan<- int,3)
	readChan<-10
	 //num:=<-readChan //error
	// fmt.Println(num)
	
	//只读的chan
	var writeChan  <-chan int
	writeChan = make(<-chan int,3)
	num1:=<-writeChan
	//writeChan<-10//error
	fmt.Println(num1)
}
