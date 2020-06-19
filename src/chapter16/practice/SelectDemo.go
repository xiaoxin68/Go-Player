package main

import (
	"fmt"
)

//selet解决从管道中取数据阻塞问题
func main() {
	intChan := make(chan int, 10)
	go func() {
		for i := 1; i <= 10; i++ {
			intChan <- i
		}
	}()

	strChan := make(chan string, 5)
	go func() {
		for i := 1; i <= 5; i++ {
			strChan <- "hello" + fmt.Sprintf("%d", i)
		}
	}()

	//传统方法在遍历管道时，不关闭会阻塞导致deadlock

	//实际开发中不好确定什么时候关闭管道，可以用select解决
	//label:
	for {
		select {
		//这里如果intChan一直没有关闭，不会一直阻塞而deadlock，会自动到下一个case
		case v := <-intChan:
			fmt.Printf("从intChan读取的数据%d\n", v)
		case v := <-strChan:
			fmt.Printf("从strChan读取的数据%s\n", v)
		default:
			fmt.Printf("都取不到了，不玩了，程序员可以加入逻辑\n")
			//break label
			return
		}
	}
}
