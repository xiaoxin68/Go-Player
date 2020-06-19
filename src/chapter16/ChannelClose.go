package main

import "fmt"

func main() {
	intChan := make(chan int, 10)
	intChan <- 10
	intChan <- 20
	close(intChan) //使用内置函数close可以关闭channel，当channel关闭后，就不能再向channel写数据了，但是仍然可以从该channel读取数据
	//intChan<-30  //panic: send on closed channel
	
	n1 := <-intChan
	fmt.Println(n1)
}
