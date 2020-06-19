package main

import "fmt"

type Dog struct {
	Name string
	Age  int
}

func main() {
	var ptrChan chan *Dog
	ptrChan = make(chan *Dog, 10)
	
	dog1 := Dog{
		Name: "张三",
		Age:  20,
	}
	dog2 := Dog{
		Name: "李四",
		Age:  30,
	}
	
	ptrChan <- &dog1
	ptrChan <- &dog2
	
	//取出
	fmt.Printf("ptrChan[1]=%v,ptrChan[2]=%v\n", *(<-ptrChan), *(<-ptrChan))
	
	//输出：ptrChan[1]={张三 20},ptrChan[2]={李四 30}
}
