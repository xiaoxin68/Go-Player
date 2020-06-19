package main

import "fmt"

func main() {
	intChan := make(chan int, 100)
	for i:=1;i<=100;i++ {
		intChan<-i
	}
	
	//遍历管道不能使用普通的for循环
	
	close(intChan)
	
	//在遍历时，如果channel没关闭，则会出现deadlock的错误:fatal error: all goroutines are asleep - deadlock!
	//在遍历时，如果channel已经关闭，则会正常遍历数据，遍历完后，就会退出遍历
	for v:=range intChan {
		fmt.Println("v=",v)
	}
}
