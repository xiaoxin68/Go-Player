package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type integer int

func main() {
	var interChan chan interface{}
	interChan = make(chan interface{}, 10)
	
	n1 := Person{
		Name: "啦啦",
		Age:  20,
	}
	
	n2 := integer(10)
	
	interChan <- n1
	interChan <- n2
	interChan <- 23.89
	interChan <- false
	interChan <- "jack"
	
	//t := <-interChan
	//fmt.Println(t.(Person).Name)
	
	fmt.Printf("interChan[1]=%v,interChan[2]=%v,interChan[3]=%v,interChan[5]=%v,interChan[5]=%v\n", <-interChan, <-interChan, <-interChan, <-interChan, <-interChan)
	
	//输出：interChan[1]={啦啦 20},interChan[2]=10,interChan[3]=23.89,interChan[5]=false,interChan[5]=jack
}
