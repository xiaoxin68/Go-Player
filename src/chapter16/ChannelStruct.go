package main

import "fmt"

type Cat struct {
	Name string
	Age int
}
func main() {
	var structChan chan Cat
	structChan = make(chan Cat,10)
	 
	 cat1 := Cat{
		 Name: "张三",
		 Age:  20,
	 }
	 cat2 := Cat{
		 Name: "李四",
		 Age:  30,
	 }
	
	structChan<-cat1
	structChan<-cat2
	
	//取出
	fmt.Printf("structChan[1]=%v,structChan[2]=%v\n", <-structChan, <-structChan)
	
	//输出：structChan[1]={张三 20},structChan[2]={李四 30}
}
